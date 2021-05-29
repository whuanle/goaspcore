package ioc

import (
	"fmt"
	"log"
	"reflect"
)

type ServiceProvider struct {
	descriptors []ServiceDescriptor
}

// 获取对象
func (s *ServiceProvider) GetService(serviceType reflect.Type) interface{} {
	index, error := s.containIndex(serviceType)
	if error != nil {
		panic(error)
	}
	descriptor := s.descriptors[index]
	if descriptor.Lifetime == Transient {
		// 实例化当前类型
		obj := descriptor.InitHandler()
		// 检查当前结构体是否还有需要被注入的字段
		obj = s.CreateObject(obj)
		return obj
	}
	return &descriptor.ImplementationInstance
}

// 查找对象的位置
func (s *ServiceProvider) containIndex(serviceType reflect.Type) (int, error) {
	for i, sd := range s.descriptors {
		if sd.ServiceType == serviceType {
			return i, nil
		}
	}
	return 0, fmt.Errorf("没有找到需要的对象，%t", serviceType)
}

// 递归给需要依赖注入的结构体字段注入实例
func (s *ServiceProvider) CreateObject(obj interface{}) interface{} {
	t := reflect.TypeOf(&obj)
	v := reflect.ValueOf(&obj).Elem()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("注入的类型不是结构体！")
		return nil
	}
	fieldNum := t.NumField()
	// 找到需要被依赖注入的字段
	for i := 0; i < fieldNum; i++ {
		tag := t.Field(i).Tag.Get("injection")
		if tag == "" {
			continue
		}
		if tag == "true" {
			value := s.GetService(t.Field(i).Type)
			v.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return &obj
}
