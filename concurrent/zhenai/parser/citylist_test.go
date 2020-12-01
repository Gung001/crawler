package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")
	const size = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Requests) != size {
		t.Errorf("The result's Requests should be %d, but it was %d", size, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url %s,but waw %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != size {
		t.Errorf("The result's Items should be %d, but it was %d", size, len(result.Items))
	}
}
