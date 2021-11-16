package router

// StaticRouter 静态路由
type StaticRouter struct {
	Url string

	Controller IControllerBase

	Action ActionDescriptor
}
