package worker

import "google/crawler/engine"

type CrawlService struct{}

func (CrawlService) Process(
	req Request, result *ParseResult) error { //序列化Request; 序列化ParseResult
	engineReq, err := DeserializeRequest(req) //将Request反序列化为引擎可用的engine.Request
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq) //调用引擎的Worker函数执行实际爬取和解析
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult) //序列化
	return nil
}
