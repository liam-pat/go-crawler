package parser

import (
	"fmt"
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
	"strings"
	"time"
)

var basicInfoRex = regexp.MustCompile(`"basicInfo":\[([^]]+)],`)
var genderRex = regexp.MustCompile(`"genderString":"([^"]+)"`)
var nickNameRex = regexp.MustCompile(`"nickname":"([^"]+)"`)
var registeredRex = regexp.MustCompile(`"objectWorkCityString":"([^"]+)"`)
var educationRex = regexp.MustCompile(`"educationString":"([^"]+)"`)
var incomeRex = regexp.MustCompile(`"月收入:([^"]+)"`)
//var imageRex = regexp.MustCompile(`"photoURL":"([^"]+)"`)
var imageRex2 = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\(([^?]+\.jpg)\?[^)]+\);`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	time.Sleep(time.Second * 10)
	profile := model.Profile{}

	profile, err := decodeBasicInfo(extractString(contents, basicInfoRex), profile)

	if err != nil {
		result := engine.ParseResult{
			Items: []interface{}{profile},
		}
		return result
	}

	profile.Gender = extractString(contents, genderRex)
	profile.Name = extractString(contents, nickNameRex)
	profile.Registered = extractString(contents, registeredRex)
	profile.Education = extractString(contents, educationRex)
	profile.Income = extractString(contents, incomeRex)
	profile.ImageUrl = extractString(contents, imageRex2)

	result := engine.ParseResult{
		Items: []interface{}{profile},
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
	profile.Marriage = arr[0]
	profile.Age = arr[1]
	profile.Height = arr[3]

	return profile, nil
}
