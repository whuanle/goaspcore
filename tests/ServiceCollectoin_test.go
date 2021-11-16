package tests

import (
	"fmt"
	"github.com/whuanle/goaspcore/ioc"
	"reflect"
	"testing"
)

type serviceB struct {
	A int
}

func (s *serviceB) Printf() {
	fmt.Println("测试")
}

func TestServiceCollection_AddScoped(t *testing.T) {
	var s ioc.IServiceCollection = &ioc.ServiceCollection{}
	s.AddScopedHandler(reflect.TypeOf(serviceB{}), func() interface{} {
		return &serviceB{
			A: 666,
		}
	})
	provider := s.Build()
	b := provider.GetService(reflect.TypeOf(serviceB{}))

	obj, ok := (*b).(*serviceB)
	if !ok {
		fmt.Println("接口转为结构体失败")
	}
	obj.Printf()
	fmt.Println(obj.A)
}
