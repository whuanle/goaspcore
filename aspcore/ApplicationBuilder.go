package aspcore

import (
	"aspcore/Middleware"
	"fmt"
	"reflect"
)

type ApplicationBuilder struct {
	middlewares map[reflect.Type]Middleware.IMiddleware
}

// 加入中间件到管道
func (builder *ApplicationBuilder) AddMiddleware(middleware Middleware.IMiddleware) {
	if middleware == nil {
		fmt.Errorf("%t 中间件不能为空")
	}
	if builder.middlewares == nil {
		builder.middlewares = map[reflect.Type]Middleware.IMiddleware{}
	}

	midType := reflect.TypeOf(middleware)
	builder.middlewares[midType] = middleware
}

func (builder *ApplicationBuilder) CopyTo() []Middleware.IMiddleware {
	mids := make([]Middleware.IMiddleware, len(builder.middlewares))
	i := 0
	for _, mid := range builder.middlewares {
		mids[i] = mid
		i++
	}
	return mids
}
