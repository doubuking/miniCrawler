package view

import (
	"testing"
	"text/template"
	"miniCrawler/frontend/model"
	"os"
)

func TestTemplate(t *testing.T)  {
	//从文件打开html文件
	template := template.Must(template.ParseFiles("template.html"))

	page :=model.SearchResult{}

	//创建一个文件
	out,err := os.Create("template.test.html")

	if err != nil{
		panic(err)
	}

	//输出到屏幕
	//err := template.Execute(os.Stdout,page)


	//输出到文件
	err = template.Execute(out,page)
	if err != nil{
		panic(err)
	}
}
