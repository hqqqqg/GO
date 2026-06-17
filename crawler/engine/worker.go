package engine

import (
	"google/crawler/fetcher"
	"log"
)

// 将parser和fetch合起来
func worker(
	r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)  //打印日志，在抓谁
	body, err := fetcher.Fetch(r.Url) //用url抓源码存到body
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s:%v",
			r.Url, err)
		return ParseResult{}, err //ParseResult是结构体
	}
	return r.ParserFunc(body, r.Url), nil //调用解析函数
}
