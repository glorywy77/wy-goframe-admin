package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AliEcsPageReq struct {
	g.Meta `path:"/api/aliEcs/page" method:"get" summary:"分页获取阿里云服务器" tags:"AliEcsService"`

	RegionId string
	// PrivateIpAddresses string
	// PublicIpAddresses  string
	// InstanceName       string
	DynamicSelectKey   string
	DynamicSelectValue string
	PageSize           int `d:"10" v:"max:100#分页数量最大100条" dc:"每页数量,最大100"`
	CurrentPage        int `d:"1"  v:"min:0#分页号码错误"     dc:"分页号码,默认1"`
}

type AliEcsPageRes struct {
	Total       int `dc:"总数"`
	Instance    []map[string]interface{}
	PageSize    int `d:"10" v:"max:100#分页数量最大100条" dc:"每页数量,最大100"`
	CurrentPage int `d:"1"  v:"min:0#分页号码错误"     dc:"分页号码,默认1"`
}
