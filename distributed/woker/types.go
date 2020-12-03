package woker

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/zhenai/parser"
	"crawler.com/distributed/config"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Item     []engine.Item
	Requests []Request
}

func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializedResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Item: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializedRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg : %v", p.Args)
		}

	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Item,
	}

	for _, req := range r.Requests {
		dr, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request : %v", err)
			continue
		}
		result.Requests = append(result.Requests, dr)
	}
	return result
}
