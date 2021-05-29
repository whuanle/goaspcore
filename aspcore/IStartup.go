package aspcore

import "ioc"

type IStartup interface {
	// 配置依赖注入
	ConfigureServices(services ioc.IServiceCollection)
	// 配置中间件
	Configure(app IApplicationBuilder)
}
