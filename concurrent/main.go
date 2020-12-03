package main

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/persist"
	"crawler.com/concurrent/scheduler"
	"crawler.com/concurrent/zhenai/parser"
	"crawler.com/distributed/config"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
const urlShanghai = "http://localhost:8080/mock/www.zhenai.com/zhenghun/suining2"

func main() {

	itemChan, err := persist.ItemSaver("dating_profile7")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:       urlShanghai,
	//	ParserFunc: parser.ParseCity,
	//})
}
