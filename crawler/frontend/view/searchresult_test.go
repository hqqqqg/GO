package view

import (
	"os"
	"testing"

	"google/crawler/engine"
	"google/crawler/frontend/view/model"
	common "google/crawler/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		t.Fatalf("create file error: %v", err)
	}
	defer out.Close()

	// 构造一份假的搜索结果
	page := model.SearchResult{}
	page.Hits = 123 // 模板会显示 "找到了 123 个相亲对象"
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: common.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Hokou:      "山东菏泽",
			Xinzuo:     "牡羊座",
			Marriage:   "离异",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
