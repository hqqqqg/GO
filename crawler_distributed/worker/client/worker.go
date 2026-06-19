package client

import (
	"google/crawler/engine"
	"google/crawler_distributed/config"
	"google/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(
	clientChan chan *rpc.Client) engine.Processor {

	return func(
		req engine.Request) (
		engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <-clientChan //从通道拿一个过来再call
		err := c.Call(config.CrawlServiceRpc,
			sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
