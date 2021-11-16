package main

import (
	"github.com/whuanle/goaspcore/aspcore"
	"github.com/whuanle/goaspcore/aspcore/router"
)

// HomeController 主页
type homeController struct {
	// 控制器的字段将会被依赖注入
	// 例如你想获得上下文
	context aspcore.HttpContext `"injection":"true"`
}

type Info struct {
}

// Init 在程序启动时会调用一次，绑定所有路由
func (controller *homeController) Init(router *router.RouterBuilder) {
	router.Post("Test", controller.Test)
	router.PostAction(controller.Test)
}

// Test 对于结构体参数，会尝试从 json 中反序列化，暂时不支持 表单，
// 对于 int string 等类型，会从 query 中取出
func (controller *homeController) Test(info Info, name string) string{
	return "11111111111111111111111111"
}
