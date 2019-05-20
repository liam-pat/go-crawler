package parser

import (
	"go-crawler/engine"
	"regexp"
)

const cityRex = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRex)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	/**
	aTag[0] the a link
	aTag[1] the first ()
	aTag[2] the second ()
	*/
	for _, aTag := range matches {
		name := aTag[2]
		result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(result.Requests, engine.Request{Url: string(aTag[1]), ParserFunc: func(bytes []byte) engine.ParseResult {
			return ParseProfile(bytes, string(name))
		}})
	}
	return result
}
