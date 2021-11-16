package router

import (
	"github.com/whuanle/goaspcore/aspcore"
	"reflect"
	"strings"
)

// RouterBuilder 路由构建器
type RouterBuilder struct {
	// 路由表
	RouterTable map[string]ActionDescriptor
}

func (router *RouterBuilder) add(static bool, url string, actionName string, requestType aspcore.RequestType) {
	router.RouterTable[url] = ActionDescriptor{
		Static:     static,
		Url:        url,
		HeaderType: requestType,
		ActionName: actionName,
	}
}

func (router *RouterBuilder) Get(url string, action interface{}) {
	static := false
	if strings.HasPrefix(url, "/") {
		static = true
	}
	router.add(static, url, reflect.TypeOf(action).Name(), aspcore.GET)
}

func (router *RouterBuilder) GetAction(action interface{}) {
	name := reflect.TypeOf(action).Name()
	router.add(false, name, name, aspcore.GET)
}

func (router *RouterBuilder) Post(url string, action interface{}) {
	static := false
	if strings.HasPrefix(url, "/") {
		static = true
	}
	router.add(static, url, reflect.TypeOf(action).Name(), aspcore.POST)
}

func (router *RouterBuilder) PostAction(action interface{}) {
	name := reflect.TypeOf(action).Name()
	router.add(false, name, name, aspcore.POST)
}

func (router *RouterBuilder) Delete(url string, action interface{}) {
	static := false
	if strings.HasPrefix(url, "/") {
		static = true
	}
	router.add(static, url, reflect.TypeOf(action).Name(), aspcore.DELETE)
}

func (router *RouterBuilder) DeleteAction(action interface{}) {
	name := reflect.TypeOf(action).Name()
	router.add(false, name, name, aspcore.DELETE)
}

func (router *RouterBuilder) Put(url string, action interface{}) {
	static := false
	if strings.HasPrefix(url, "/") {
		static = true
	}
	router.add(static, url, reflect.TypeOf(action).Name(), aspcore.PUT)
}

func (router *RouterBuilder) PutAction(action interface{}) {
	name := reflect.TypeOf(action).Name()
	router.add(false, name, name, aspcore.PUT)
}
