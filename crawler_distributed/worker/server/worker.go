package main //RPC服务器监听端口，接收远程爬取请求

import (
	"flag"
	"fmt"
	"google/crawler_distributed/rpcsupport"
	"google/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() { // 启动Worker RPC服务
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port), //拼接监听地址 port会变
		worker.CrawlService{}))    //注册CrawlService为 RPC 服务
}
