package main //RPC服务器监听端口，接收远程爬取请求

import (
	"fmt"
	"google/crawler_distributed/config"
	"google/crawler_distributed/rpcsupport"
	"google/crawler_distributed/worker"
	"log"
)

func main() { // 启动Worker RPC服务
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0), //拼接监听地址
		worker.CrawlService{}))                 //注册CrawlService为 RPC 服务
}
