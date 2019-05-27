package parser

import (
	"go-crawler/engine"
	"regexp"
)

var regexString = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(content []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}

	matches := regexString.FindAllSubmatch(content, -1)
	for _, aTag := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:    string(aTag[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}
	return result
}
