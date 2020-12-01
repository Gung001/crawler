package main

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/scheduler"
	"crawler.com/concurrent/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
const urlShanghai = "http://localhost:8080/mock/www.zhenai.com/zhenghun/suining2"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:       urlShanghai,
	//	ParseFunc: parser.ParseCity,
	//})
}
