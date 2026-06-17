package persist //启动docker的es

import (
	"encoding/json"
	"google/crawler/engine"
	"google/crawler/model"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestSave(t *testing.T) {
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
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = save(client, index, &expected)
	if err != nil {
		t.Fatalf("save error: %v", err)
	}

	resp, err := client.Get(index, expected.Id)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	t.Logf("%s", resp)

	var raw struct {
		Source json.RawMessage `json:"_source"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		t.Fatalf("decode error: %v", err)
	}

	var actual engine.Item
	if err := json.Unmarshal(raw.Source, &actual); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		t.Fatalf("FromJsonObj error: %v", err)
	}
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v;expected %v",
			actual, expected)
	}
}
