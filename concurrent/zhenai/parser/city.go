package parser

import (
	"crawler.com/concurrent/engine"
	"crawler.com/distributed/config"
	"regexp"
)

const cityReg = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/suining2/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    url,
				Parser: NewProfileParser(string(m[2])),
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			},
		)
	}

	return result
}
