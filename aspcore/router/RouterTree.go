package router

import "sync"

// RouterTree 路由树
type RouterTree struct {

	// 控制器路由
	ControllerRouter sync.Map

	// 静态路由
	StaticRouter sync.Map
}
