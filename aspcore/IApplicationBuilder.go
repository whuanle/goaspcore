package aspcore

// 构建中间件
type IApplicationBuilder interface {
	// 加入中间件
	AddMiddleware(middleware IMiddleware)
}
