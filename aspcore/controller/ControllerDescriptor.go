package controller

import (
	"github.com/whuanle/goaspcore/aspcore/router"
)

// ControllerDescriptor 控制器描述
type ControllerDescriptor struct {

	// 控制器地址
	Url string
	// 控制器
	Controller router.IControllerBase

	ActionRouters []router.ActionDescriptor
}
