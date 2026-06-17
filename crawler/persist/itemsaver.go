package persist // 负责把爬下来的数据保存到 elasticsearch

import (
	"bytes"                 // 把字节数组转成 io.Reader，给 es 客户端用
	"encoding/json"         // 把 struct 序列化成 JSON 字节流
	"errors"                // 创建一个带消息的 error
	"google/crawler/engine" // 引入 engine 包，定义了 Item 类型
	"log"

	"github.com/elastic/go-elasticsearch/v8" // 官方 elasticsearch 客户端 v8
)

func ItemSaver(index string) (chan engine.Item, error) { //
	client, err := elasticsearch.NewDefaultClient() // 创建默认 es 客户端（连本机 localhost:9200）
	if err != nil {                                 // 创建失败时
		return nil, err // 错误向上抛
	}
	out := make(chan engine.Item) // 无缓冲的item管道
	go func() {                   // goroutine 后台持续消费 out 管道
		itemCount := 0 // 已经处理了多少条 item
		for {          // 持续接收并保存 item，直到进程退出
			item := <-out
			log.Printf("Item Saver:got item "+
				"#%d:%v", itemCount, item) //要存的内容先打印
			itemCount++

			err := save(client, index, &item) // 调用 save 真正把 item 写进 elasticsearch
			if err != nil {
				log.Printf("Item Saver:error"+
					"Saving item %v:%v",
					item, err)
			}
		}
	}()
	return out, nil
}

func save(
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
