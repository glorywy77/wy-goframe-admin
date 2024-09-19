package ecs

import (
	"context"
	"errors"
	"fmt"
	"os"

	"strings"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"

	// console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"

	// "github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sEcs struct{}
)

func New() *sEcs {
	return &sEcs{}
}

func init() {
	service.RegisterEcs(New())
}

func CreateClient() (_result *ecs20140526.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。

	ak := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	sk := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")

	// 检查环境变量是否为空
	if ak == "" || sk == "" {
		return nil, errors.New("environment variables ALIBABA_CLOUD_ACCESS_KEY_ID or ALIBABA_CLOUD_ACCESS_KEY_SECRET are not set")
	}

	config := &openapi.Config{
		AccessKeyId:     tea.String(ak),
		AccessKeySecret: tea.String(sk),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Ecs
	config.Endpoint = tea.String("ecs.cn-zhangjiakou.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

func (s *sEcs) EcsPage(ctx context.Context, in model.AliEcsPageInput) (out g.Map, err error) {

	client, err := CreateClient()
	if err != nil {
		return nil, err
	}

	//此闭包函数是用于修改数据格式，以符合提交阿里云接口的字符串数组格式
	change := func(k string, v string) *string {
		if v == "" {
			return nil
		}
		if gregex.IsMatchString(`(.+),(.+)`, v) || gregex.IsMatchString(`es`, k) {

			parts := strings.Split(v, ",")
			x := []string{}
			x = append(x, parts...)
			return tea.String(gconv.String(x))
		}
		return tea.String(v)
	}

	req := g.Map{
		"RegionId":          tea.String(in.RegionId),
		in.DynamicSelectKey: change(in.DynamicSelectKey, in.DynamicSelectValue),
		"PageNumber":        tea.Int32(int32(in.CurrentPage)),
		"PageSize":          tea.Int32(int32(in.PageSize)),
	}

	var describeInstancesRequest *ecs20140526.DescribeInstancesRequest
	gconv.Struct(req, &describeInstancesRequest)
	runtime := &util.RuntimeOptions{}

	// 使用 defer 和 recover 捕获运行时错误
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("请求阿里云API失败,请检查AccessKey,Recovered from panic: %v", r)
			g.Log().Errorf(ctx, errMsg)
			err = fmt.Errorf(errMsg)
		}
	}()

	resp, err := client.DescribeInstancesWithOptions(describeInstancesRequest, runtime)
	if err != nil {
		g.Log().Errorf(ctx, "Error calling DescribeInstancesWithOptions: %v", err)
		return nil, err
	}
	out = gconv.Map(resp)
	return
}
