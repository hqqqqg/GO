// 正则表达式
package main

import (
	"fmt"
	"regexp"
)

// const text = `
// My email is ccmouse@gmail.com@abc.com
// `
const text = `
My email is ccmouse@gmail.com@abc.com
email is abc@def.org
email2 is  kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	// re := regexp.MustCompile(".+@.+\\..+") //获取正则表达 .代表任何字符，若要代表实际的点，用\\.
	// re := regexp.MustCompile(`.+@.+\..+`) //单引号好懂，一定要match个点
	// match := re.FindString(text) //找到符合正则表达式的子串

	// re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`) //[a-zA-Z0-9]这里不包括空格
	// re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	// match := re.FindAllString(text, -1) //传个-1代表要找所有的匹配

	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1) //返回一个二维slice
	for _, m := range match {
		fmt.Println(m)
	}
	fmt.Println(match)
}
