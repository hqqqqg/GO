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
		name := string(m[2])
		result.Items = append(
			result.Items, "User"+name) //将用户名字作为item返回出去
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(
					c []byte) engine.ParseResult {
					return ParseProfile( //用户信息
						c, name)
				},
			})
	}
	return result

}
