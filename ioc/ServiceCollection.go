package ioc

import (
	"fmt"
	"log"
	"reflect"
)

// ServiceCollection 注册对象
type ServiceCollection struct {
	// 容器内的对象
	descriptors []ServiceDescriptor
	// 对象数量
	Count int
}

// 基础注入方法
func (s *ServiceCollection) addAny(implementationType reflect.Type, serviceType reflect.Type, lifetime ServiceLifetime, f func() interface{}) {
	descriptor := ServiceDescriptor{
		Name:          serviceType.Name(),
		InheritType:   implementationType,
		ImplementType: serviceType,
		Lifetime:      lifetime,
		InitHandler:   f,
	}
	// 单例模式提前实例化，也就是常驻内存
	if lifetime == Singleton {
		descriptor.ImplementationInstance = f()
	}

	s.add(descriptor)
}

func (s *ServiceCollection) AddScoped(t reflect.Type) {
	f := func() interface{} {
		return reflect.New(t).Elem().Interface()
	}
	s.addAny(t, t, Scope, f)
}

func (s *ServiceCollection) AddScopedHandler(t reflect.Type, f func() interface{}) {
	s.addAny(t, t, Scope, f)
}

func (s *ServiceCollection) AddScopedForm(implementationType reflect.Type, serviceType reflect.Type) {
	f := func() interface{} {
		return reflect.New(serviceType).Elem().Interface()
	}
	s.addAny(implementationType, serviceType, Scope, f)
}

func (s *ServiceCollection) AddScopedHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{}) {
	if !serviceType.Implements(implementationType) {
		panic(fmt.Sprintf("%t 不实现 %t", serviceType, implementationType))
	}
	s.addAny(implementationType, serviceType, Scope, f)
}

func (s *ServiceCollection) AddSingleton(t reflect.Type) {
	f := func() interface{} {
		return reflect.New(t).Elem().Interface()
	}
	s.addAny(t, t, Singleton, f)
}

func (s *ServiceCollection) AddSingletonHandler(t reflect.Type, f func() interface{}) {
	s.addAny(t, t, Singleton, f)
}

func (s *ServiceCollection) AddSingletonForm(implementationType reflect.Type, serviceType reflect.Type) {
	f := func() interface{} {
		return reflect.New(serviceType).Elem().Interface()
	}
	s.addAny(implementationType, serviceType, Singleton, f)
}

func (s *ServiceCollection) AddSingletonHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{}) {
	s.addAny(implementationType, serviceType, Singleton, f)
}

func (s *ServiceCollection) AddTransient(t reflect.Type) {
	f := func() interface{} {
		return reflect.New(t).Elem().Interface()
	}
	s.addAny(t, t, Transient, f)
}

func (s *ServiceCollection) AddTransientHandler(t reflect.Type, f func() interface{}) {
	s.addAny(t, t, Transient, f)
}

func (s *ServiceCollection) AddTransientForm(implementationType reflect.Type, serviceType reflect.Type) {
	f := func() interface{} {
		return reflect.New(serviceType).Elem().Interface()
	}
	s.addAny(implementationType, serviceType, Transient, f)
}

func (s *ServiceCollection) AddTransientHandlerForm(implementationType reflect.Type, serviceType reflect.Type, f func() interface{}) {
	s.addAny(implementationType, serviceType, Transient, f)
}

func (s *ServiceCollection) add(serviceDescriptor ServiceDescriptor) {

	if s.descriptors == nil {
		s.descriptors = make([]ServiceDescriptor, 0, 8)
	}

	s.descriptors = append(s.descriptors, serviceDescriptor)
	s.Count = len(s.descriptors)
}

// 移除
func (s *ServiceCollection) remove(descriptor ServiceDescriptor) {
	index, err := s.containIndex(descriptor)
	if err != nil {
		log.Fatal(err)
		return
	}
	s.descriptors = append(s.descriptors[:index], s.descriptors[index:]...)
	s.Count = len(s.descriptors)
}

// 查找对象的位置
func (s *ServiceCollection) containIndex(descriptor ServiceDescriptor) (int, error) {
	for i, sd := range s.descriptors {
		if sd.ImplementType == descriptor.ImplementType {
			return i, nil
		}
	}
	return 0, fmt.Errorf("没有找到对于的对象，%v", descriptor)
}

func (s *ServiceCollection) Build() IServiceProvider {
	// scoped
	descriptors := make([]ServiceDescriptor, s.Count, s.Count)

	// 复制集合中的所有对象到新的容器中，并且对每个 Scope 的对象实例化
	for i, descriptor := range s.descriptors {
		descriptors[i] = descriptor
		if descriptors[i].Lifetime == Scope {
			descriptors[i].ImplementationInstance = descriptors[i].InitHandler()
		}
	}

	var services IServiceProvider
	services = &ServiceProvider{
		descriptors: descriptors,
	}
	return services
}

func (s *ServiceCollection) CopyTo() IServiceCollection {
	descriptors := make([]ServiceDescriptor, s.Count, s.Count)

	for i, descriptor := range s.descriptors {
		descriptors[i] = descriptor
	}
	return &ServiceCollection{
		descriptors: descriptors,
		Count:       len(descriptors),
	}
}
