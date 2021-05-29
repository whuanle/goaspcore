package ioc

import "reflect"

// 依赖注入提供器
type IServiceProvider interface {
	// 获取你需要的服务实例
	GetService(serviceType reflect.Type) interface{}
}
