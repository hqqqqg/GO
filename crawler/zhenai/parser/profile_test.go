// 测试个人资料解析器
package parser

import (
	"os"
	"testing"

	"google/crawler/engine"
	"google/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := os.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parseProfile(contents,
		"http://album.zhenai.com/u/108906739", "安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Item should contain 1"+"element; but was %v", result.Items)
	}
	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	if actual != expected {
		t.Errorf("expect %v; but was %v",
			expected, actual)
	}
}
