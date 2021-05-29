package Middleware

import "aspcore"

type Middleware struct {
}

func (m *Middleware) Invoke(context aspcore.HttpContext) {

	// 执行你的管道方法
	// 执行下一个中间件
	context.Next()

	// 继续执行你的管道方法
}
