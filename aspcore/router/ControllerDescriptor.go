package router

import (
	"aspcore/Controllers"
)

type ControllerDescriptor struct {

	// 控制器地址
	Url string

	// 控制器
	Controller Controllers.ControllerBase

	ActionRouters []ActionDescriptor
}
