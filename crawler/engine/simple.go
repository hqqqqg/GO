// 获得seed，维护一个request队列，对每一个request去fetch，结果放在body里面
package engine

import (
	"log"
)

// 简单调度器
type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) { //可以丢任意数量的种子
	var requests []Request //任务队列
	for _, r := range seeds {
		requests = append(requests, r) //种子塞进排队通道
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,
			parseResult.Requests...) //...citylist的结果展开一个一个加进去，
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
