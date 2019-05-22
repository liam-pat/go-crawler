package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	PayLoad interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
