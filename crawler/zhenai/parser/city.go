package parser

import (
	"google/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(https?://[^.]*\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>` //album.zhenai.com

func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1) //所有匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(
			result.Items, "User"+string(m[2])) //将城市名字作为item返回出去
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result

}
