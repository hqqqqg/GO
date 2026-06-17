package parser //

import ( // 导入依赖
	"google/crawler/engine" // 有 ParseResult、Item 等公共结构
	"google/crawler/model"  // Profile 结构体
	"regexp"                // 做正则匹配，从 HTML 里抠字段
	"strconv"               // strconv.Atoi 字符串转成整数
)

var genderRe = regexp.MustCompile(`<td><span class="label">性别:</span>([^<]+)</td>`)                                          // 预编译性别正则，捕获括号里是性别值
var ageRe = regexp.MustCompile(`<td><span class="label">年龄:</span>([^<]+)</td>`)                                             // 改回提取字符串 // 年龄正则
var heightRe = regexp.MustCompile(`<td><span class="label">身高:</span>([^<]+)</td>`)                                          // 身高正则
var weightRe = regexp.MustCompile(`<td><span class="label">体重:</span>([^<]+)</td>`)                                          // 体重正则
var incomeRe = regexp.MustCompile(`<td><span class="label">月薪:</span>([^<]+)</td>`)                                          // 月薪正则
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况:</span>([^<]+)</td>`)                                        // 婚况正则
var educationRe = regexp.MustCompile(`<td><span class="label">学历:</span>([^<]+)</td>`)                                       // 学历正则
var occupationRe = regexp.MustCompile(`<td><span class="label">职业:</span>([^<]+)</td>`)                                      // 职业正则
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯:</span>([^<]+)</td>`)                                           // 籍贯正则
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座:</span>([^<]+)</td>`)                                          // 星座正则
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件:</span>([^<]+)</td>`)                                         // 住房条件正则
var carRe = regexp.MustCompile(`<td><span class="label">是否购车:</span>([^<]+)</td>`)                                           // 是否购车正则
var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album\.zhenai\.com/u/[\d]+)"[^>]*>([^<]+)</a>`) // 猜你喜欢区域的推荐用户，第 1 组是 URL，第 2 组是名字
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)                                                        // 从用户主页 URL 里抽出纯数字 ID

func ParseProfile( // 给 engine 用的解析函数，解析一个用户主页
	contents []byte, url string,
	name string) engine.ParseResult { //
	profile := model.Profile{}
	profile.Name = name
	// 流水线式装填数据
	//注意前面几个
	age, err := strconv.Atoi( // Atoi 转成整数
		extractString(contents, ageRe)) // 负责返回正则捕获到的内容
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
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
		Items: []engine.Item{ // 解析出的数据条目列表
			{
				Url:  url,
				Type: "zhenai", // 业务类型标识，es 存储时要
				Id: extractString( // 从 URL 里抠出唯一 ID
					[]byte(url), idUrlRe),
				Payload: profile, // 资料数据放在 Payload 里
			},
		},
	}
	matches := guessRe.FindAllSubmatch( // 在 HTML 里查找所有“猜你喜欢”用户
		contents, -1) // 找全
	for _, m := range matches { // 遍历每一个匹配
		result.Requests = append(result.Requests, // 把这些推荐用户作为新任务追加到结果
			engine.Request{ // 构造一条新的 Request 交给 engine 调度
				Url:        string(m[1]), // 推荐用户的主页 URL
				ParserFunc: ProfileParser(string(m[2])),
			})
	}
	return result // 返回 engine
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents) // 找第一个匹配，返回 [][]byte，match[0] 是整段，match[1] 是第一个括号!!
	if len(match) >= 2 {
		return string(match[1]) // 转成 string 返回
	} else { // 没匹配上
		return ""
	}
}

func ProfileParser(
	name string) engine.ParserFunc {
	return func(
		c []byte, url string) engine.ParseResult { // 参数 c 是下载好的 HTML
		return ParseProfile(c, url, name) // 复用同一个解析器

	}
}
