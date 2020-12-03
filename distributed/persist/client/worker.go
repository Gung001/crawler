package client

import (
	"crawler.com/concurrent/engine"
	"crawler.com/distributed/config"
	"crawler.com/distributed/woker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := woker.SerializedRequest(req)
		var sResult woker.ParseResult
		c := <-clientChan
		err := c.Call(config.ItemCrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return woker.DeserializeResult(sResult), nil
	}
}
