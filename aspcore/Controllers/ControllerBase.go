package Controllers

import "aspcore/router"

type ControllerBase interface {
	Init(router *router.RouterBuilder)
}
