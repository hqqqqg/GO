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
	contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(
		contents, -1) //所有匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(
			result.Requests, engine.Request{
				Url: url,
				ParserFunc: func(
					c []byte) engine.ParseResult {
					return ParseProfile( //用户信息
						c, url, name)
				},
			})
	}
	matches = cityUrlRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result

}
