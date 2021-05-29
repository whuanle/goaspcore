package main

import (
	"aspcore"
	"aspcore/Controllers"
	"ioc"
	"reflect"
)

type StartUp struct {
}

// 配置依赖注入
func (s *StartUp) ConfigureServices(services ioc.IServiceCollection) {

	// 注入 mvc 服务
	services.AddSingletonHandler(reflect.TypeOf(Controllers.ControllerService{}), func() interface{} {
		controllers := Controllers.ControllerService{}
		controllers.AddController(&aspcore.HomeController{})
		return controllers
	})
}

// 配置中间件
func (s *StartUp) Configure(app aspcore.IApplicationBuilder) {

	// 设定 mvc 中间件
	app.AddMiddleware(&Controllers.MvcMiddleware{})
}
