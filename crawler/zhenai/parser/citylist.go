package parser

import (
	"google/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(https?://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>` //有必要再深究

func ParseCityList(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) //所有匹配
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(
			result.Items, "City "+string(m[2])) //将城市名字作为item返回出去
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity, //城市里面找用户
			})
		limit--
		if limit == 0 {
			break
		}
	}
	return result

}
