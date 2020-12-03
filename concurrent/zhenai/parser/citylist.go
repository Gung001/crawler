package parser

import (
	"crawler.com/concurrent/engine"
	"crawler.com/distributed/config"
	"regexp"
)

const cityListReg = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListReg)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}

	return result
}
