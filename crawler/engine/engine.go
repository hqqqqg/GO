// 获得seed，维护一个request队列，对每一个request去fetch，结果放在body里面
package engine

import (
	"google/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) { //可以丢任意数量的种子
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

// 将parser和fetch合起来
func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)  //打印日志，在抓谁
	body, err := fetcher.Fetch(r.Url) //抓源码存到body
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s:%v",
			r.Url, err)
		return ParseResult{}, err //ParseResult是结构体
	}
	return r.ParserFunc(body), nil
}
