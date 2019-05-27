package engine

import (
	fetcher "go-crawler/fetch"
	"log"
)

func Worker(firstRequest Request) (ParseResult, error) {

	log.Printf("Fetching %s", firstRequest.Url)
	body, err := fetcher.Fetch(firstRequest.Url)

	if err != nil {
		log.Printf("Fetch:error "+"fetch url %s : %v", firstRequest.Url, err)

		return ParseResult{}, err
	}

	ParseResult := firstRequest.Parser.Parse(body, firstRequest.Url)

	return ParseResult, nil
}
