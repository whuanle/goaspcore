package main

import (
	"fmt"
	"github.com/whuanle/goaspcore/aspcore"
)

// MyMiddleware 中间件实例
type MyMiddleware struct {
}

func (m *MyMiddleware) Invoke(context aspcore.HttpContext) {

	request := context.Request
	fmt.Println(request.RequestURI)

	// 执行你的管道方法
	// 执行下一个中间件
	context.Next()

	response := context.Response
	fmt.Println(len(response.Header()))
	// 继续执行你的管道方法
}
