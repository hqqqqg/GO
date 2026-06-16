package persist

import (
	"context"
	"encoding/json"
	"google/crawler/model"
	"io"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestSave(t *testing.T) {
	testData := model.Profile{
		Name:       "测试账号",
		Gender:     "男",
		Age:        28,
		Height:     175,
		Weight:     70,
		Income:     "30-50万",
		Marriage:   "未婚",
		Education:  "本科",
		Occupation: "软件工程师",
		Hokou:      "北京",
		Xinzuo:     "天蝎座",
		House:      "已购房",
		Car:        "已购车",
	}

	id, err := save(testData)
	if err != nil {
		panic(err)
	}
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	resp, err := client.Get(
		"dating_profile",
		id,
		client.Get.WithContext(context.Background()), //上下文参数
	)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	t.Logf("%s", string(bodyBytes))

	var wrapper struct {
		Source model.Profile `json:"_source"`
	}

	// 把全部的 JSON 字节流解析到这个结构体里，自动填入键值对
	err = json.Unmarshal(bodyBytes, &wrapper)
	if err != nil {
		panic(err)
	}

	// 把剥出来的真实数据赋值给 actual
	actual := wrapper.Source

	// 最后进行比对
	if actual != testData {
		t.Errorf("got %v;expected %v", actual, testData)
	}
}
