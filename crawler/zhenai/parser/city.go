package parser

import (
	"google/crawler/engine"
	"regexp"
)

// 先编译快
var (
	profileRe = regexp.MustCompile(
		`<a href="(https?://[^.]*\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`) //album.zhenai.com
	cityUrlRe = regexp.MustCompile(
		`href = "http://[^.]*\.zhenai.com/zhenghun/shanghai/[^"]+"`) //匹配下一页
)

func ParseCity(
	contents []byte, url string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(
		contents, -1) //所有匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				Parser: NewProfileParser(
					string(m[2])),
			})
	}
	matches = cityUrlRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, "ParseCity"),
			})
	}
	return result

}
