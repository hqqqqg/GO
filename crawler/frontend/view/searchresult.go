//把 model.SearchResult 数据套进 template.html

package view

import (
	"google/crawler/frontend/view/model"
	"io"
	"text/template"
)

type SearchResultView struct {
	tmpl *template.Template //解析好的模板对象
}

func CreateSearchResultView(
	filename string) SearchResultView {
	return SearchResultView{
		tmpl: template.Must(
			template.ParseFiles(filename)), //调用 ParseFiles 解析模板
	}
}

func (s SearchResultView) Render( //把数据 data 套进模板
	w io.Writer, data model.SearchResult) error {
	return s.tmpl.Execute(w, data)
}
