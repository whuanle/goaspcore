package aspcore

import (
	"reflect"
)

type ApplicationBuilder struct {
	middlewares map[reflect.Type]IMiddleware
}

// AddMiddleware 加入中间件到管道
func (builder *ApplicationBuilder) AddMiddleware(mid IMiddleware) {
	if mid == nil {
	panic("%t 中间件不能为空")
	}
	if builder.middlewares == nil {
		builder.middlewares = map[reflect.Type]IMiddleware{}
	}

	midType := reflect.TypeOf(mid)
	builder.middlewares[midType] = mid
}

func (builder *ApplicationBuilder) CopyTo() []IMiddleware {
	mids := make([]IMiddleware, len(builder.middlewares))
	i := 0
	for _, mid := range builder.middlewares {
		mids[i] = mid
		i++
	}
	return mids
}
