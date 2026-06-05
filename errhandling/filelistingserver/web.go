// 通过网页读文件内容,出错处理，尽量不要用panic 用error
// fib可读， fib2不可
package main

import (
	"google/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, //向浏览器发送
	request *http.Request) error //用户发来请求

// 出错处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) { //输入是函数输出也是函数
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			// log.Warn("Error handling request: %s", err.Error)
			log.Printf("Error occurred"+"handling request:%s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest) //将userErr.Message()给用户看，并附上错误类型
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound //404
			case os.IsPermission(err):
				code = http.StatusForbidden //403
			default:
				code = http.StatusInternalServerError //500
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error            //给系统看
	Message() string //给用户看
}

func main() {
	http.HandleFunc("/", //以这个开头的都交给后面的函数处理
		errWrapper(filelisting.HandleFileList)) //包起来
	err := http.ListenAndServe(":8888", nil) //启动服务器，nill代表使用默认配置的路由表
	if err != nil {
		panic(err)
	}
}
