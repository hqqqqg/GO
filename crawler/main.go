// 爬虫相亲网 用ai重新改了代码，现在的网址已经从gbk改成UTF-8了
// go get golang.org/x/text 安装gbk.go改编码为UTF-8s 位置在/go/pkg/mod/golang.org/x/text@v0.23.0/encoding/simplifiedchinese/gbk.go
// go get golang.org/x/net/html/charset
package main

import (
	"google/crawler/engine"
	"google/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
