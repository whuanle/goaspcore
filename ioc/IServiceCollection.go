package ioc

import "reflect"

// 对象容器
type IServiceCollection interface {

	// 域实例注册
	AddScoped(t reflect.Type)
	// 实例注册，并自定义如何初始化实例
	AddScopedHandler(t reflect.Type, f func() interface{})
	// 实例注册，注册接口及其实现
	AddScopedForm(implementationType reflect.Type, serviceType reflect.Type)
	// 实例注册，注册接口及其实现，并自定义如何初始化实例
	AddScopedHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{})

	// 单例注册
	AddSingleton(t reflect.Type)
	AddSingletonHandler(t reflect.Type, f func() interface{})
	AddSingletonForm(implementationType reflect.Type, serviceType reflect.Type)
	AddSingletonHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{})

	// 瞬时实例注册
	AddTransient(t reflect.Type)
	AddTransientHandler(t reflect.Type, f func() interface{})
	AddTransientForm(implementationType reflect.Type, serviceType reflect.Type)
	AddTransientHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{})

	// 复制当前容器的所有对象
	CopyTo() IServiceCollection
	// 构建 ioc 服务
	Build() IServiceProvider
}
