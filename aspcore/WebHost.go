package aspcore

import (
	"fmt"
	"github.com/whuanle/goaspcore/ioc"
	"net/http"
	"reflect"
)

type WebHost struct {
}

func (host *WebHost) Run(start *IStartup, addr string) {
	fmt.Println("正在初始化")
	web := &Host{}
	web.Init(start)
	fmt.Println("正在监控")
	err := http.ListenAndServe(addr, web)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

type Host struct {
	// 全局，在 Startup 中构建完成的依赖注入
	serviceCollection *ioc.IServiceCollection

	// 启动器
	startup *IStartup

	// 中间件
	middlewares []IMiddleware
}

func (host *Host) Init(startup *IStartup) {
	// 注入所有依赖
	var serviceCollection ioc.IServiceCollection = &ioc.ServiceCollection{}
	(*startup).ConfigureServices(serviceCollection)
	host.serviceCollection = &serviceCollection
	host.startup = startup

	// 获取中间件
	var app IApplicationBuilder
	app = &ApplicationBuilder{}
	(*host.startup).Configure(app)
	host.middlewares = (app).(*ApplicationBuilder).CopyTo()

	// 中间件也要注入到容器中
	for _, mid := range host.middlewares {
		serviceCollection.AddScoped(reflect.TypeOf(mid))
	}

}

// 处理所有请求输出
func (host *Host) ServeHTTP(response http.ResponseWriter, resuest *http.Request) {
	// 为每个访问的用户串口一个依赖注入容器
	// 注入一些有用的对象，默认注入
	serviceCollection := (*host.serviceCollection).CopyTo()
	serviceCollection.AddScopedHandler(reflect.TypeOf(resuest), func() interface{} {
		return resuest
	})

	// 构建上下文
	context := &HttpContext{
		Request:     resuest,
		Response:    response,
		middlewares: host.middlewares,
		index:       0,
		// ...
	}
	defer func() {
		code := int(context.StatusCode)
		if code == 0 {
			context.StatusCode = Status404NotFound
			_, _ = fmt.Fprint(context.Response, "goaspcore 框架提醒你，404 页面")
		}
	}()

	// 调用中间件
	context.Next()

}
