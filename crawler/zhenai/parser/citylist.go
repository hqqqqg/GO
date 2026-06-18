package parser

import (
	"google/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(https?://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>` //有必要再深究

func ParseCityList(
	contents []byte, url string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) //所有匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, "ParseCity"),
			})
	}
	return result

}
