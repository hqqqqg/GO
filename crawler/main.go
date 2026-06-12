// 爬虫相亲网 用ai重新改了代码，现在的网址已经从gbk改成UTF-8了
// go get golang.org/x/text 安装gbk.go改编码为UTF-8s 位置在/go/pkg/mod/golang.org/x/text@v0.23.0/encoding/simplifiedchinese/gbk.go
// go get golang.org/x/net/html/charset
package main

import (
	"bufio"    //缓冲I/O提供偷看数据但不拿走的功能
	"fmt"      //格式化输出
	"io"       //基础I/O
	"net/http" //网络通讯 发送HTTP请求

	"golang.org/x/net/html/charset"      //官方扩展包 自动嗅探HTML网页的编码
	"golang.org/x/text/encoding"         //文本编码接口
	"golang.org/x/text/encoding/unicode" //Unicode编码
	"golang.org/x/text/transform"        //文本转换器 把别的编码翻译成GO原生的UTF-8
)

func main() {
	// 1. 伪装自己：不要直接用 http.Get，而是自己构建一个 Request
	request, err := http.NewRequest(http.MethodGet, "http://www.zhenai.com/zhenghun", nil)
	if err != nil {
		panic(err) //网址格式不对，直接报错
	}
	// 加上真实浏览器的 User-Agent，骗过反爬虫保安
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	// 2. 发送请求
	client := http.Client{}         //实例化一个客户机
	resp, err := client.Do(request) //让客户机带着伪装的request访问网站
	if err != nil {
		panic(err) //断网或者目标网站挂了
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { //200
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// 3. 智能转码：让程序自己猜网页是什么编码
	// 我们需要用 bufio 包装一下，因为猜编码需要试读前面的一小段字节，bufio 允许读完再塞回去
	bufReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bufReader) //判断是gbk还是utf-8

	// 4. 根据猜出来的编码（e）进行转换（如果是 UTF8，底层就不会做多余的转换）
	utf8Reader := transform.NewReader(bufReader, e.NewDecoder())
	all, err := io.ReadAll(utf8Reader) //读数据
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all) //转换为字符串并打印
}

// 这是一个极其好用的通用函数，专门用来“智能探测”网页编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 试着偷看前 1024 个字节，不改变读取指针
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("Fetcher error: %v\n", err)
		return unicode.UTF8 // 如果偷看失败，默认当做 UTF-8 处理
	}
	// 引入 golang.org/x/net/html/charset 库来自动判断编码
	//DetermineEncoding分析这1024个字节里面的HTML标签 比如<meta charset="gbk">
	//返回这个网页的真实编码e
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
