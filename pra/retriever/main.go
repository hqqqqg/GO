package main

import (
	"fmt"
	"pra/retriever/mock"
	"pra/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http.//www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute}
	inspect(r)

	if mockretriever, ok := r.(mock.Retriever); ok {
		fmt.Println("contents:", mockretriever.Contents) //老是大小写混着写
	} else {
		fmt.Println("not a mock retriever")
	}

}

func inspect(r Retriever) {
	fmt.Printf("%T,%v\n", r, r) //接口里面是类型和值/指针
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

}
