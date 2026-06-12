// 接口

package main

import (
	"fmt"
	"google/retriever/mock"
	"google/retriever/real"
	"time"
	//引入
)

type Retriever interface {
	Get(url string) string //定义接口，接口里有一个方法Get，参数是url，返回值是string
}
type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "1",
			"course": "2",
		})
}

// 接口组合，实现包含的接口的功能
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "2",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute}
	inspect(r)

	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	// fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r) //接口变量里面有类型和指针/值
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
