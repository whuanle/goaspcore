package router

import "github.com/whuanle/goaspcore/aspcore"

// ActionDescriptor  action 描述
type ActionDescriptor struct {
	// 是否是静态的
	Static bool
	// 路由地址，url
	Url string
	// 请求类型
	HeaderType aspcore.RequestType
	// action 名称
	ActionName string
}
