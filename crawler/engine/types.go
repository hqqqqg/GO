package engine

type ParserFunc func(
	content []byte, url string) ParseResult
type Request struct {
	Url        string
	ParserFunc ParserFunc //下载完源码后调用ParserFunc函数解析
}

type ParseResult struct {
	Requests []Request //任务队列
	Items    []Item    //抠出来的城市名，用户的信息
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
