package api

//  通用的请求/响应结构体

// 通用的分页请求/响应结构体
type CommonPaginationReq struct {
	PageSize int `d:"10" v:"max:1000#分页数量最大1000条" dc:"每页数量,最大1000"  json:"pageSize"`
	CurrentPage int `d:"1"  v:"min:0#分页号码错误"     dc:"分页号码,默认1"  json:"currenPage"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数" json:"total"`
}



type CommonResultRes struct {
	Result string `json:"result"`
}