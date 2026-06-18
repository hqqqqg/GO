// 先起server/main里面的服务器
package main

import (
	"fmt"
	"google/crawler/engine"
	"google/crawler/scheduler"
	"google/crawler/zhenai/parser"
	"google/crawler_distributed/config"
	"google/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort)) //让远程的服务器去save
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		// Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
