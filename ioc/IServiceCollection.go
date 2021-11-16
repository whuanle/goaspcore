package ioc

import "reflect"

// IServiceCollection 是依赖注入对象容器接口
type IServiceCollection interface {

	// AddScoped 注册对象为域实例，即 Scope 生命周期
	AddScoped(t reflect.Type)

	// AddScopedHandler 注册对象为域实例，即 Scope 生命周期，并自定义如何初始化实例
	AddScopedHandler(t reflect.Type, f func() interface{})
	// AddScopedForm 注册对象为域实例，注册接口及其实现
	AddScopedForm(implementationType reflect.Type, serviceType reflect.Type)
	// AddScopedHandlerForm 注册对象为域实例，注册接口及其实现，并自定义如何初始化实例
	AddScopedHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{})

	// AddSingleton 注册对象为单例
	AddSingleton(t reflect.Type)
	// AddSingletonHandler 注册对象为域实例，即 Scope 生命周期，并自定义如何初始化实例
	AddSingletonHandler(t reflect.Type, f func() interface{})
	// AddSingletonForm 注册对象为域实例，注册接口及其实现
	AddSingletonForm(implementationType reflect.Type, serviceType reflect.Type)
	// AddSingletonHandlerForm 注册对象为域实例，注册接口及其实现，并自定义如何初始化实例
	AddSingletonHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{})

	// AddTransient 注册对象为单例
	AddTransient(t reflect.Type)
	// AddTransientHandler 注册对象为域实例，即 Scope 生命周期，并自定义如何初始化实例
	AddTransientHandler(t reflect.Type, f func() interface{})
	// AddTransientForm 注册对象为域实例，注册接口及其实现
	AddTransientForm(implementationType reflect.Type, serviceType reflect.Type)
	// AddTransientHandlerForm 注册对象为域实例，注册接口及其实现，并自定义如何初始化实例
	AddTransientHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{})

	// CopyTo 复制当前容器的所有对象，生成新的容器
	CopyTo() IServiceCollection
	// 	Build() 构建依赖注入服务提供器 IServiceProvider
	Build() IServiceProvider
}
