package fetcher

import (
	"bufio"
	"fmt"
	"go-crawler/crawler_distributed/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// this you can control the crawler fetch speed
var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	log.Printf("fetch url : %v", url)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// sometime this site have the tech to gen the cookie in js
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3100.0 Safari/537.36")
	req.Header.Set("Cookie", "Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1558348617; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1558348585; FSSBBIl1UgzbN7N80T=3y_FEXfOIhumpbxk3uwKo3EXg1Af4T_gw9NVrR6CPL9Yfq3uZZweLWxgztxlcu6hj14lOUBIU2zh1iS9DGWHX_ahpP4R2swtH2b2LLVGfkeylRkD55oxaEVDfwSb1cpHZA03uPMF24GxSGDCXBxcfZkyN_o5.zkDtrkbnXKEPZ_v3YsI7rHAL.NbMDC02p5MnrOlE50MJvTckQthl.BIPLk0JRyNl0oE.GEZhzncUHFwSIRtyQx2CraVDFO7tDyLFjrouiDEb4ff12kLz5cudWjxsoUA0EVJTdCMPlO.urIEP3Ms5hP5F4Hi_Fsxr865uJzV; sid=3af15b39-d841-4d1b-8e43-89ac14e72760; FSSBBIl1UgzbN7N80S=9jVXOwn6ThqTV.1sm6xh3ra3ghbWNlzISMhKF4XFGGZzEPrzLurTqcHWFLbWe63H")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong code with the request,code status %d,message: %s", response.StatusCode, response.Status)
	}

	// tran the html to charset utf-8
	bodyReader := bufio.NewReader(response.Body)
	encode := getDetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// golang have some fun to get the content 's charset
func getDetermineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)

	if err != nil {
		log.Printf("Fetch error : %v", err)
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")

	return encode
}
