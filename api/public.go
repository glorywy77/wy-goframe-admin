package api

//  通用的请求/响应结构体

// 通用的分页请求/响应结构体
type CommonPaginationReq struct {
	PageSize    int `d:"10" v:"max:100#分页数量最大100条" dc:"每页数量,最大100"  json:"pageSize"`
	CurrentPage int `d:"1"  v:"min:0#分页号码错误"     dc:"分页号码,默认1"  json:"currenPage"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数" json:"total"`
}

type CommonResultRes struct {
	Result string `json:"result"`
}

// // 阿里云通用的分页请求/响应结构体
// type AliCommonPageReq struct {
// 	PageSize   int32 `d:"10" v:"max:100#分页数量最大100条" dc:"每页数量,最大100"  json:"pageSize"`
// 	PageNumber int32 `d:"1"  v:"min:0#分页号码错误"     dc:"分页号码,默认1"  json:"pageNumber"`
// }

// type AliCommonPageRes struct {
// 	TotalCount int32 `dc:"总数" json:"totalCount"`
// }
