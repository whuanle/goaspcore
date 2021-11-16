package ioc

import "reflect"

// IServiceProvider 依赖注入提供器
type IServiceProvider interface {
	// GetService 获取你需要的服务实例
	GetService(serviceType reflect.Type) *interface{}
}
