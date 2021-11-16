package main

import (
	"fmt"
	"github.com/whuanle/goaspcore/ioc"
	"log"
	"reflect"
)

type IAnimal interface {
	Println(s string)
}
type Dog struct {
}

func (my Dog) Println(s string) {
	fmt.Println(s)
}

func Example() {
	// ** 正常代码 **
	// 接口
	var animal IAnimal
	imy := reflect.TypeOf(&animal).Elem()
	// 结构体
	my := reflect.TypeOf(Dog{})

	// ** 依赖注入容器 **

	// 创建容器
	var collection ioc.IServiceCollection = &ioc.ServiceCollection{}

	// 注入服务，生命周期为 scoped
	collection.AddScopedForm(imy, my)

	// 构建服务 Provider
	serviceProvider := collection.Build()

	// 获取对象
	// *interface{}
	obj := serviceProvider.GetService(my)

	// *interface{} -> Dog
	log.Print(reflect.TypeOf(obj))
	log.Print(reflect.TypeOf((*obj).(Dog)))
}

func main() {
	Example()
}
