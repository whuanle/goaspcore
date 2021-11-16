package controller

import (
	"github.com/whuanle/goaspcore/aspcore/router"
	"reflect"
)

// ControllerService 控制器注册服务
type ControllerService struct {
	controllers map[string]router.IControllerBase
}

// AddController 注入新的控制器
func (c *ControllerService) AddController(controller router.IControllerBase) {
	if c.controllers == nil {
		c.controllers = map[string]router.IControllerBase{}
	}
	c.controllers[reflect.TypeOf(controller).Name()] = controller
}
