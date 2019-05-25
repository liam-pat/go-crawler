package parser

import (
	"fmt"
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
	"strings"
)

var basicInfoRex = regexp.MustCompile(`"basicInfo":\[([^]]+)],`)
var genderRex = regexp.MustCompile(`"genderString":"([^"]+)"`)
var nickNameRex = regexp.MustCompile(`"nickname":"([^"]+)"`)
var registeredRex = regexp.MustCompile(`"objectWorkCityString":"([^"]+)"`)
var educationRex = regexp.MustCompile(`"educationString":"([^"]+)"`)
var incomeRex = regexp.MustCompile(`"月收入:([^"]+)"`)
//var imageRex = regexp.MustCompile(`"photoURL":"([^"]+)"`)
var imageRex2 = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\(([^?]+\.jpg)\?[^)]+\);`)
var idUrlRex = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}

	profile, err := decodeBasicInfo(extractString(contents, basicInfoRex), profile)

	if err != nil {
		result := engine.ParseResult{}
		return result
	}

	profile.Gender = extractString(contents, genderRex)
	profile.Name = extractString(contents, nickNameRex)
	profile.Name = name
	profile.Registered = extractString(contents, registeredRex)
	profile.Education = extractString(contents, educationRex)
	profile.Income = extractString(contents, incomeRex)
	profile.ImageUrl = extractString(contents, imageRex2)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "true_love",
				Id:      extractString([]byte(url), idUrlRex),
				PayLoad: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) > 1 {
		return string(match[1])
	}
	return ""
}

func decodeBasicInfo(jsonString string, profile model.Profile) (model.Profile, error) {
	if jsonString == "" {
		return profile, fmt.Errorf("User hidden her information! %s ", jsonString)
	}
	arr := strings.Split(string(jsonString), ",")
	profile.Marriage = trimQuotes(string(arr[0]))
	profile.Age = trimQuotes(string(arr[1]))
	profile.Height = trimQuotes(string(arr[3]))

	return profile, nil
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
