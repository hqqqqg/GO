package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"google/crawler/engine"
	"google/crawler_distributed/config"
	"google/crawler_distributed/rpcsupport"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func ItemSaver(
	host string) (chan engine.Item, error) { //
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item) // 无缓冲的item管道
	go func() {                   // goroutine 后台持续消费 out 管道
		itemCount := 0 // 已经处理了多少条 item
		for {          // 持续接收并保存 item，直到进程退出
			item := <-out
			log.Printf("Item Saver:got item "+
				"#%d:%v", itemCount, item) //要存的内容先打印
			itemCount++
			//call rpc to save item
			result := ""
			err := client.Call(
				config.ItemSaverRpc,
				item, &result)
			if err != nil {
				log.Printf("Item Saver:error"+
					"Saving item %v:%v",
					item, err)
			}
		}
	}()
	return out, nil
}

func Save(
	client *elasticsearch.Client, index string, //由外面告诉index
	item *engine.Item) error {
	if item.Type == "" { // item 必须有 Type
		return errors.New("must supply Type") // 返回错误
	}

	body, err := json.Marshal(item) // 把 item 序列化成 JSON 字节数组
	if err != nil {
		return err // 返回错误
	}

	resp, err := client.Index(index, // 向 dating_profile 索引写入
		bytes.NewReader(body),                // 用 Reader 包装一下交给 es 客户端
		client.Index.WithDocumentID(item.Id)) // 用 item.Id 作为 id
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return errors.New(resp.String())
	}
	return nil // 表示成功
}
