package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong code with the request,code status %d,message: %s", response.StatusCode,response.Status)
	}

	bodyReader := bufio.NewReader(response.Body)
	encode := getDetermineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func getDetermineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)

	if err != nil {
		log.Printf("Fetch error : %v", err)
		return unicode.UTF8
	}

	encode, _, _ := charset.DetermineEncoding(bytes, "")

	return encode
}
