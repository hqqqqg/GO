package persist

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item "+
				"#%d:%v", itemCount, item) //要存的内容先打印
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) (
	id string, err error) {
	client, err := elasticsearch.NewDefaultClient() //初始化，不用手动关闭sniff
	//Must turn off sniff in docker
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(item) //item转换为json字节，body只接受字节流
	resp, err := client.Index(
		"dating_profile",      // 索引名称,http://localhost:9200/dating_profile/_search可搜索到
		bytes.NewReader(data), // 文档内容 ，包装成 Reader，
		client.Index.WithContext(context.Background()),
	)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 定义一个匿名结构体，专门用来接收 JSON 里的 "_id" 字段
	var result struct {
		Id string `json:"_id"`
	}

	// 将 resp.Body (字节流) 解析到我们的结构体中
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Id, nil
}
