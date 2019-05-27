package parser

import (
	"go-crawler/engine"
	"regexp"
)

var profileRex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var nextPage = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[\d]+)">下一页</a>`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRex.FindAllSubmatch(contents, -1)
	/**
	aTag[0] the a tag  eg. <a></a>
	aTag[1] the first () eg. url
	aTag[2] the second () eg. name
	*/
	for _, aTag := range matches {
		name := aTag[2]
		// because the result.items should not to save ,so set empty
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:    string(aTag[1]),
				Parser: NewProfileParser(string(name)),
			})
	}

	// to get the next page
	matches = nextPage.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}

	return result
}
