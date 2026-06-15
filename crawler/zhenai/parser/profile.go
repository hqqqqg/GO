package parser

import (
	"google/crawler/engine"
	"google/crawler/model"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<th><span class="label">姓名:</span>([^<]+)</td>`) // 名字通常在别的地方，比如 <h1> 里，请根据实际情况修改
var genderRe = regexp.MustCompile(`<td><span class="label">性别:</span>([^<]+)</td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄:</span>([^<]+)</td>`) // 改回提取字符串
var heightRe = regexp.MustCompile(`<td><span class="label">身高:</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重:</span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月薪:</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况:</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历:</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业:</span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯:</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座:</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件:</span>([^<]+)</td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车:</span>([^<]+)</td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	// 流水线式装填数据
	//注意前面几个
	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Name = extractString(contents, nameRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	// 把装满数据的档案袋放进
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1]) // 返回括号里捕获到的干净数据
	} else {
		return ""
	}
}
