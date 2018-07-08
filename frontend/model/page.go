package model

type SearchResult struct {
	Hits int64 //总条数
	Start int	//开始页
	Query string //搜索的参数
	PrevFrom int //上一页
	NextFrom int //下一页
	Items []interface{} //查询的数据
}
