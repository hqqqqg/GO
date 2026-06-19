// 分布式爬虫 启动ItemSaver RPC客户端
package main

import (
	"flag"
	"google/crawler/engine"
	"google/crawler/scheduler"
	"google/crawler/zhenai/parser"
	"google/crawler_distributed/config"
	itemsaver "google/crawler_distributed/persist/client" //RPC客户端
	"google/crawler_distributed/rpcsupport"
	worker "google/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver( //启动客户端
		*itemSaverHost) //服务地址
	if err != nil {
		panic(err)
	}
	pool := createClientPool(
		strings.Split(*workerHosts, ",")) //分割

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{ // 创建并发引擎
		Scheduler: &scheduler.QueuedScheduler{},
		// Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList, //使用城市列表解析器
	// })
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), //使用城市页解析器解析该页面的用户列表
	})
}

func createClientPool(
	hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf(
				"error connecting to %s:%v",
				h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
