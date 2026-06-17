// 测试城市解析器
package parser

import (
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// contents, err := fetcher.Fetch(
	// 	"http://www.zhenai.com/zhenghun")
	// if err != nil {
	// 	panic(err)
	// }
	// err = os.WriteFile("citylist_test_data.html", contents, 0644) //存文件

	contents, err := os.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 10 //城市个数
	//取三个
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/baicheng1",
		"http://www.zhenai.com/zhenghun/cangzhou",
	}
	// expectedCities := []string{
	// 	"City 阿坝", "City 白城", "City 沧州",
	// }

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d:%s;but"+"was %s",
				i, url, result.Requests[i].Url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+
			"requests; but had %d",
			resultSize, len(result.Items))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected city #%d:%s;but"+"was %s",
				i, url, result.Requests[i].Url)
		}
	}

}
