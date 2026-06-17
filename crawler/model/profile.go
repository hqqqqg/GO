package model

import "encoding/json"

type Profile struct {
	Url        string
	Id         string
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o) //打包成字节
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile) //安装profile装填
	return profile, err
}
