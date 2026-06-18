package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	"google/crawler/frontend/view"       //  把数据套进模板
	"google/crawler/frontend/view/model" // 数据模型 SearchResult
)

type SearchResultHandler struct { ////直接喂给 http.Handle
	view   view.SearchResultView // 把 template.html 解析进内存的模板
	client *elasticsearch.Client //连接本地的 elasticsearch 客户端 9200
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elasticsearch.NewDefaultClient() //连本地 ES
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(
			template), // 把模板路径交给 view 解析
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP( //http.Handler 接口的实现
	w http.ResponseWriter, //作为 HTTP 响应返回浏览器
	req *http.Request) { //查询
	// FormValue 会自动解析
	q := strings.TrimSpace(req.FormValue("q"))

	// from 是偏移量
	from, err := strconv.Atoi(
		req.FormValue("from"))
	if err != nil { // 转换失败就当成第 1 页(from=0)
		from = 0
	}

	// 拿着 q 和 from 去问 ES, 返回数据
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}
	// 把 page 套进 template.html, 写回浏览器
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	const (
		index    = "dating_profile"
		pageSize = 10 // 每页显示 10 条。

	)
	var result model.SearchResult // 查 ES, 把结果装进 model.SearchResult。
	result.Query = q
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(map[string]interface{}{
		"query": map[string]interface{}{ //{"query": {"query_string": {"query": "q 本身"}}}
			"query_string": map[string]interface{}{
				// 直接用 q, 不 rewrite
				"query": q,
			},
		},
	}); err != nil {
		return result, err
	}

	resp, err := h.client.Search(
		h.client.Search.WithContext(context.Background()),
		h.client.Search.WithIndex(index), //限定索引
		h.client.Search.WithBody(&body),
		//分页
		h.client.Search.WithFrom(from),
		h.client.Search.WithSize(pageSize),
	)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return result, fmt.Errorf("es search error: %s", resp.String())
	}

	var raw struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"` // 用于显示 "找到了 N 个"
			} `json:"total"`
			Hits []struct {
				Source interface{} `json:"_source"` //_source结构:{"Url": "...", "Type": "...", "Id": "...", "Payload": {...}}

			} `json:"hits"`
		} `json:"hits"`
	}
	//ES 返回的 JSON 响应体解析到 raw
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return result, err
	}

	result.Hits = raw.Hits.Total.Value
	result.Start = from
	// 塞进 Items
	for _, hit := range raw.Hits.Hits {
		result.Items = append(result.Items, hit.Source)
	}

	// 简单分页
	result.PrevFrom = result.Start - pageSize
	// 在第 1 页时 PrevFrom 会算成 -10, 截断到 0
	if result.PrevFrom < 0 {
		result.PrevFrom = 0
	}
	// 下一页直接 +pageSize
	result.NextFrom = result.Start + pageSize

	return result, nil
}

// 不用每次都写payload
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:") //$1就是第一个括号括起来的部分
}
