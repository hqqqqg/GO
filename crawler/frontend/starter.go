// 独立的 HTTP 服务

package main

import (
	"google/crawler/frontend/controller"
	"net/http"
)

func main() {
	// 用 http.Dir 是为了限制在指定目录里
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view")))

	// 把模板文件 template.html 解析进内存
	http.Handle("/search", //   /search交给 controller.SearchResultHandler 处理
		controller.CreateSearchResultHandler(
			"crawler/frontend/view/template.html"))

	// 监听  nil 表示用默认的路由表
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
