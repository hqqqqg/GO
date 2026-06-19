// 分布式爬虫 启动ItemSaver RPC客户端
package main

import (
	"fmt"
	"google/crawler/engine"
	"google/crawler/scheduler"
	"google/crawler/zhenai/parser"
	"google/crawler_distributed/config"
	itemsaver "google/crawler_distributed/persist/client" //RPC客户端
	worker "google/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver( //启动客户端
		fmt.Sprintf(":%d", config.ItemSaverPort)) //服务地址
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
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
