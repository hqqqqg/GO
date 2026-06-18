package engine

type ParserFunc func(
	content []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser //下载完源码后调用ParserFunc函数解析
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
type NilParser struct {
}

func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (
	name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(
	contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}
func (f *FuncParser) Serialize() (
	name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
