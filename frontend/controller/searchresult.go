package controller

import (
	"miniCrawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"miniCrawler/frontend/model"
	"context"
	"reflect"
	"miniCrawler/engine"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

//初始化
func CreateSearchResultHandler(template string)SearchResultHandler  {
	client,err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil{
		panic(err)
	}

	return SearchResultHandler{
		view:view.CreateSearchResultView(template),
		client:client,
	}
}


//localhost:8888/search?q=  &from=
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	//获取相关参数
	q := strings.TrimSpace(req.FormValue("q"))

	from,err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w,"q=%s,from=%d",q,from)

	var page model.SearchResult
	page,err = h.getSearchResult(q,from)

	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

	err = h.view.Render(w,page)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

}

//根据条件获取数据并返回
func (h SearchResultHandler)getSearchResult(q string,from int) (model.SearchResult,error) {
	var result model.SearchResult
	//搜索条件
	result.Query = q

	//查询操作
	resp,err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).Do(context.Background())

	if err != nil{
		return result,err
	}
	//总条数
	result.Hits = resp.TotalHits()
	//开始条数
	result.Start = from

	//查询的数据
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	//上一页
	result.PrevFrom = result.Start - len(result.Items)
	//下一页
	result.NextFrom = result.Start + len(result.Items)

	return result,nil
}

