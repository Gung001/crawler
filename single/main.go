package main

import (
	"crawler.com/single/engine"
	"crawler.com/single/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCityList,
	})
}
