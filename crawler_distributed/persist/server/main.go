// 启动RPC服务器监听端口，接收远程保存请求
package main

import (
	"fmt"
	"google/crawler_distributed/config"
	"google/crawler_distributed/persist"
	"google/crawler_distributed/rpcsupport" //RPC框架工具
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config.ItemSaverPort), //拼接监听地址
		config.ElasticIndex))                     //使用配置中指定的ES索引名
}

func serveRpc(host, index string) error { // 启动 RPC 服务
	client, err := elasticsearch.NewDefaultClient() // 创建默认ES客户端连接 :9200
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
