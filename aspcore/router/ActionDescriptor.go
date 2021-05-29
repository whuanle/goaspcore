package router

import "aspcore"

// Action 描述
type ActionDescriptor struct {
	// 是否是静态的
	Static bool
	// url
	Url string
	// 请求类型
	HeaderType aspcore.RequestType
	// action 名称
	ActionName string
}
