package parser

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "心事痕迹迁就", "http://localhost:8080/mock/album.zhenai.com/u/180900910")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element")
	}

	profile := result.Items[0]

	_profile := model.Profile{
		Name:       "心事痕迹迁就",
		Age:        39,
		Gender:     "男",
		Height:     185,
		Weight:     200,
		Income:     "8001-10000元",
		Marriage:   "离异",
		Education:  "大学",
		Occupation: "程序员",
		Hokou:      "西安市",
		Xinzuo:     "双子座",
		House:      "租房",
		Car:        "无车",
	}

	expected := engine.Item{
		Id:      "180900910",
		Url:     "http://localhost:8080/mock/album.zhenai.com/u/180900910",
		Type:    "zhenai",
		Payload: _profile,
	}

	if profile != expected {
		t.Errorf("expected %v ,but was %v", expected, profile)
	}

}
