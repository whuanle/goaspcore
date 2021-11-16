package aspcore

import "github.com/whuanle/goaspcore/ioc"

type IStartup interface {
	// ConfigureServices 配置依赖注入
	ConfigureServices(services ioc.IServiceCollection)
	// Configure 配置中间件
	Configure(app IApplicationBuilder)
}
