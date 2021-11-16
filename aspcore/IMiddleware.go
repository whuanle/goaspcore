package aspcore

// IMiddleware 中间件接口
type IMiddleware interface {
	// Invoke 执行中间件
	Invoke(context *HttpContext)
}
