package Middleware

import "aspcore"

// 中间件
type IMiddleware interface {
	Invoke(context aspcore.Context)
}
