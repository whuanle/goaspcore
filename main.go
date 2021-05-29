package main

import (
	"fmt"
	"ioc"
	"log"
	"reflect"
)

type IMy interface {
	Println(s string)
}
type My struct {
}

func (my My) Println(s string) {
	fmt.Println(s)
}



func Example() {
	// 接口
	var imytemp IMy
	imy := reflect.TypeOf(&imytemp).Elem()
	// 结构体
	my := reflect.TypeOf(My{})
	// 创建容器
	var serviceCollectoin ioc.IServiceCollection = &ioc.ServiceCollection{}

	// 注入服务，生命周期为 scoped
	serviceCollectoin.AddScopedForm(imy, my)

	// 构建服务 Provider
	serviceProvider := serviceCollectoin.Build()

	// 获取对象
	obj := serviceProvider.GetService(my)

	log.Print(reflect.TypeOf(obj))
	log.Print(reflect.TypeOf(obj.(My)))
}

func main(){
	Example()
}