package model

type SearchResult struct {
	Hits     int64  //命中率
	Start    int    //起始
	Query    string //搜索
	Items    []interface{}
	PrevFrom int
	NextFrom int
}
