package view

import (
	"html/template"
	"io"
	"miniCrawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}
//打开文件
func CreateSearchResultView(filename string) SearchResultView  {
	return SearchResultView{
		template:template.Must(template.ParseFiles(filename)),
	}
}

//输出文件
func (s SearchResultView)Render(w io.Writer,data model.SearchResult) error {
	return s.template.Execute(w,data)
}
