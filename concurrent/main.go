package main

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/persist"
	"crawler.com/concurrent/scheduler"
	"crawler.com/concurrent/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
const urlShanghai = "http://localhost:8080/mock/www.zhenai.com/zhenghun/suining2"

func main() {

	itemChan, err := persist.ItemSaver("dating_profile2")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:       urlShanghai,
	//	ParserFunc: parser.ParseCity,
	//})
}
