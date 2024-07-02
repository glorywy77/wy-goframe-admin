package api

//  通用的请求/响应结构体

// 通用的分页请求/响应结构体
type CommonPaginationReq struct {
	Size int `d:"10" v:"max:100#分页数量最大100条" dc:"每页数量,最大100"`
	Page int `d:"1"  v:"min:0#分页号码错误"     dc:"分页号码,默认1"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数"`
}
