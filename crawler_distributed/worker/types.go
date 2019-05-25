package worker

import (
	"errors"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"go-crawler/crawler_distributed/config"
	"go-crawler/engine"
	"go-crawler/sites/com.zhenai/parser"
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
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{Items: r.Items}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeSerializeRequest(r Request) (engine.Request, error) {
	parser1, err := DeSerializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser1,
	}, nil
}

func DeSerializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeSerializeRequest(req)
		if err != nil {
			log.Error("error deSerializing request: %v ", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func DeSerializeParser(p SerializedParser) (engine.Parser, error) {

	//Logger.Printf("parserName: %v", p.Name)

	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v ", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unKnown Requests Name")
	}
}
