package woker

import "crawler.com/concurrent/engine"

type CrawlService struct {
}

func (CrawlService) Process(request Request, result *ParseResult) error {
	deserializeRequest, err := DeserializeRequest(request)
	if err != nil {
		return err
	}

	parseResult, err := engine.Worker(deserializeRequest)
	if err != nil {
		return err
	}

	*result = SerializedResult(parseResult)
	return nil
}
