package fetcher

import (
	"bufio"
	"fmt"
	"google/crawler_distributed/config"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(
	time.Second / config.Qps) //每10毫秒发送一次，防止网站限流

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetch url %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { //200
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	//丢进determineEncoding里面的事带缓冲带的
	e := determineEncoding(resp.Body) //判断是gbk还是utf-8
	// 4. 根据猜出来的编码（e）进行转换（如果是 UTF8，底层就不会做多余的转换）
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return io.ReadAll(utf8Reader) //读数据

}

func determineEncoding(r io.Reader) encoding.Encoding {
	// 试着偷看前 1024 个字节，不改变读取指针
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8 //返回默认的encoding
	}
	// 引入 golang.org/x/net/html/charset 库来自动判断编码
	//DetermineEncoding分析这1024个字节里面的HTML标签 比如<meta charset="gbk">
	//返回这个网页的真实编码e
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
