package worker

import (
	"errors"
	"fmt"
	"google/crawler/engine"
	"google/crawler/zhenai/parser"
	"google/crawler_distributed/config"
	"log"
)

type SerializedParser struct { //网络传输用的解析器
	Name string //用于反序列化时找到对应的解析函数
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser //序列化的解析器
}

type ParseResult struct {
	Items   []engine.Item
	Request []Request
}

func SerializeRequest(r engine.Request) Request { // 序列化
	name, args := r.Parser.Serialize() //Serialize方法获取Name和Args
	return Request{                    //构造网络传输的Request
		Url: r.Url,
		Parser: SerializedParser{ //设置序列化后的解析器
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult( //序列化ParseResult
	r engine.ParseResult) ParseResult {
	result := ParseResult{ //初始化结果结构体
		Items: r.Items,
	}
	for _, req := range r.Requests { //遍历引擎返回的每个Request
		result.Request = append(result.Request,
			SerializeRequest(req))
	}
	return result
}

func DeserializeRequest( //反序列化为引擎可用的
	r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser) //根据Parser.Name建具体解析器
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{ //构造引擎可用的Request
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult( //ParseResult反序列化
	r ParseResult) engine.ParseResult {
	result := engine.ParseResult{ // 初始化
		Items: r.Items,
	}
	for _, req := range r.Request {
		engineReq, err := DeserializeRequest(req) //反序列化
		if err != nil {
			log.Printf("error deserializing"+
				"request:%v", err)
			continue
		}
		result.Requests = append(result.Requests,
			engineReq)
	}
	return result
}

func deserializeParser( //根据解析器名称和参数重建engine.Parser
	p SerializedParser) (engine.Parser, error) {
	switch p.Name { //根据解析器名称
	case config.ParseCityList: // 城市列表解析器
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), nil
	case config.ParseCity: // 城市页解析器
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile: //用户资料解析器
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(
				userName), nil
		} else {
			return nil, fmt.Errorf("invalid"+
				"arg:%v", p.Args)
		}
	default:
		return nil, errors.New(
			"unknown parser name")
	}
}
