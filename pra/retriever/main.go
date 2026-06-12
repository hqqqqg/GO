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

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func post(p Poster) {
	p.Post(url,
		map[string]string{
			"name":   "1",
			"course": "2",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "1",
	})
	return s.Get(url)

}

func download(r Retriever) string {
	return r.Get(url)

}

func main() {
	var r Retriever
	retriever := mock.Retriever{"fake"}
	r = &retriever
	// d := download(r)
	// fmt.Println(d)
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute}
	inspect(r)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("contents:", mockRetriever.Contents)
	} else {
		fmt.Println("no ")
	}

	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspenct", r)
	fmt.Printf(">%T %v\n", r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}

}
