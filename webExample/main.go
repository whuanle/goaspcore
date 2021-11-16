package main

import (
	"github.com/whuanle/goaspcore/aspcore"
	"github.com/whuanle/goaspcore/aspcore/controller"
	"github.com/whuanle/goaspcore/ioc"
	"reflect"
)

func main() {
	host := aspcore.WebHost{}
	var start aspcore.IStartup = &Startup{}
	host.Run(&start,"localhost:8182")
}

type Startup struct {
}

// ConfigureServices 配置依赖注入
func (s *Startup) ConfigureServices(services ioc.IServiceCollection) {

	// 注入 mvc 服务
	services.AddSingletonHandler(reflect.TypeOf(controller.ControllerService{}), func() interface{} {
		controllers := controller.ControllerService{}
		controllers.AddController(&homeController{})
		return controllers
	})
}

// Configure 配置中间件
func (s *Startup) Configure(app aspcore.IApplicationBuilder) {

	// 设定 mvc 中间件
	app.AddMiddleware(&controller.MvcMiddleware{})
}
