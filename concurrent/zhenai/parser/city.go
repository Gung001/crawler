package parser

import (
	"crawler.com/concurrent/engine"
	"regexp"
)

const cityReg = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/suining2/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(
			result.Items, "User "+name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParseFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, name) // 闭包传递用户名字
				},
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			},
		)
	}

	return result
}
