package router

import "aspcore/Controllers"

// 静态路由
type StaticRouter struct {
	Url string

	Controller Controllers.ControllerBase

	Action ActionDescriptor
}
