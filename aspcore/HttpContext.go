package aspcore

import (
	"github.com/whuanle/goaspcore/ioc"
	"net/http"
	"reflect"
)

type HttpContext struct {
	// 客户端请求
	Request *http.Request

	// 服务端响应
	Response http.ResponseWriter

	// 响应代码
	StatusCode StatusCodes

	// 当前中间件的索引
	index int

	// 要执行的中间件
	middlewares []IMiddleware

	// 此次请求的 ioc 对象
	ioc ioc.IServiceProvider
}

func (context *HttpContext) GetService(t reflect.Type) interface{} {
	return context.ioc.GetService(t)
}

// Next 执行管道中的下一个中间件
func (context *HttpContext) Next() {
	if context.index < len(context.middlewares) {
		// 当前要执行的中间件
		thisMiddleware := context.middlewares[context.index]

		// 中间件不直接执行，而是通过依赖注入实例化
		mid := context.GetService(reflect.TypeOf(thisMiddleware)).(IMiddleware)
		// 执行此中间件
		mid.Invoke(context)
	}
}
