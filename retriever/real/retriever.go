package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url) //发出http请求
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true) //返回http文本，true返回body内容
	defer resp.Body.Close()                          //关闭响应的body
	if err != nil {
		panic(err)
	}
	return string(result) //把http响应的内容转成字符串
}
