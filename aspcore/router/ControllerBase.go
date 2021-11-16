package aspcore

import "github.com/whuanle/goaspcore/aspcore/router"

// IControllerBase 控制器接口
type IControllerBase interface {
	Init(router *router.RouterBuilder)
}
