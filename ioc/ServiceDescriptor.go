package ioc

import (
	"reflect"
)

// ServiceDescriptor 是注入项的描述
type ServiceDescriptor struct {
	// 对象名称
	Name string

	// 对象生命周期
	Lifetime ServiceLifetime

	// 对象继承的接口，要注入的接口等
	InheritType reflect.Type

	// 实现对象，实例对象
	ImplementType reflect.Type

	// 已被实例化的对象，存储在内存中
	ImplementationInstance interface{}

	// 如何实例化对象
	InitHandler func() interface{}
}