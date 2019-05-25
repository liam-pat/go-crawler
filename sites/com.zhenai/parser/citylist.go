package parser

import (
	"go-crawler/engine"
	"regexp"
)

const regexString = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte, url string) engine.ParseResult {

	re := regexp.MustCompile(regexString)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	/**
	aTag[0] the a link
	aTag[1] the first ()
	aTag[2] the second ()
	*/
	for _, aTag := range matches {
		//result.Items = append(result.Items, "City "+string(aTag[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(aTag[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
