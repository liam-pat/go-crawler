package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	content, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(content, "http://album.zhenai.com/u/1893685027", "朵儿")

	if len(result.Items) != 1 {
		t.Errorf("Item should hava one at lease,but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1893685027",
		Type: "true_love",
		Id:   "1893685027",
		PayLoad: model.Profile{
			Name:       "朵儿",
			Gender:     "女士",
			Age:        "23岁",   //
			Height:     "161cm", //
			Income:     "8千-1.2万",
			Marriage:   "未婚", //
			Education:  "中专",
			Registered: "上海黄浦区",
			ImageUrl:   "https://photo.zastatic.com/images/photo/14390/57559191/1506226687714.jpg",
		}}

	if actual != expected {
		t.Errorf("expect %+v,but was %+v", expected, actual)
	}
}
