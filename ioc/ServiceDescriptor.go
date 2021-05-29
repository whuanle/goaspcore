package ioc

import (
	"reflect"
)

// 对象描述
type ServiceDescriptor struct {
	// 对象名称
	Name string

	// 对象生命周期
	Lifetime ServiceLifetime

	// 对象继承的接口
	ImplementationType reflect.Type

	// 对象类型
	ServiceType  reflect.Type

	// 实例对象
	ImplementationInstance interface{}

	// 如何实例化对象
	InitHandler func() interface{}
}