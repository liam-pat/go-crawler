package parser

import (
	"go-crawler/engine"
	"regexp"
)

var profileRex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRex = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[0-9]+)"`)
var nextPage = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[\d]+)">下一页</a>`)

func ParseCity(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRex.FindAllSubmatch(contents, -1)
	/**
	aTag[0] the a link
	aTag[1] the first ()
	aTag[2] the second ()
	*/
	for _, aTag := range matches {
		name := aTag[2]
		//result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(aTag[1]),
			Parser: NewProfileParser(string(name)),
		})
	}

	//matches = cityUrlRex.FindAllSubmatch(contents, -1)
	//
	//for _, m := range matches {
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:        string(m[1]),
	//		ParserFunc: ParseCity,
	//	})
	//}

	matches = nextPage.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
