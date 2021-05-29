package Controllers

import "reflect"

type ControllerService struct {
	controllers map[string]ControllerBase
}

// 注入新的控制器
func (c *ControllerService) AddController(controller ControllerBase) {
	if c.controllers == nil {
		c.controllers = map[string]ControllerBase{}
	}
	c.controllers[reflect.TypeOf(controller).Name()] = controller
}
