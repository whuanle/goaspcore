package Controllers

import (
	"aspcore"
	"aspcore/router"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// mvc 中间件
type MvcMiddleware struct {
	controllerServices ControllerService `"injection":"true"`
	// 路由树
	RouterTree router.RouterTree

	Context aspcore.HttpContext `"injection":"true"`
}

func (m *MvcMiddleware) Invoke(context aspcore.HttpContext) {

	// 解析路由
	url := context.Request.URL.Path
	// 静态路由
	a, b := m.RouterTree.StaticRouter.Load(url)
	if b {
		router := a.(router.StaticRouter)
		controller := context.GetService(reflect.TypeOf(router.Controller))
		cb := controller.(ControllerBase)
		m.request(cb, router.Action)
	} else {
		urls := strings.Split(context.Request.URL.Path, "/")
		if len(urls) == 2 {
			controller, ok := m.RouterTree.ControllerRouter.Load(urls[0])
			if ok {
				base := controller.(router.ControllerDescriptor)

				var action router.ActionDescriptor

				for _, ac := range base.ActionRouters {
					if ac.ActionName == urls[1] {
						action = ac
						break
					}
				}
				if action.ActionName == "" {
					context.StatusCode = aspcore.Status404NotFound
					fmt.Fprint(context.Response, "goaspcore 提醒你 404")
				}
				m.request(base.Controller, action)
			}
		} else {
			context.StatusCode = aspcore.Status404NotFound
			fmt.Fprint(context.Response, "goaspcore 提醒你 404")
		}
	}

	context.Next()

}

func (m *MvcMiddleware) request(base ControllerBase, action router.ActionDescriptor) {
	controller := m.Context.GetService(reflect.TypeOf(base)).(ControllerBase)

	t := reflect.TypeOf(controller)
	personValue := reflect.ValueOf(controller)
	actionT, _ := t.MethodByName(action.ActionName)

	infoFunc := personValue.MethodByName(action.ActionName)
	paramCount := actionT.Type.NumIn()
	paramValues := make([]reflect.Value, paramCount)
	for i := 0; i < paramCount; i++ {
		paramType := actionT.Type.In(i)

		// 实例化、反序列化参数
		switch paramType.Kind() {
		case reflect.Struct:
			paramValue := reflect.New(paramType).Elem().Interface()
			m.BodyToJson(&paramValue)
			paramValues = append(paramValues, reflect.ValueOf(paramValue))
		case reflect.Int | reflect.Int32:
			paramValue, _ := strconv.Atoi(m.QueryTo(paramType.Name()))
			paramValues = append(paramValues, reflect.ValueOf(paramValue))
		case reflect.String:
			paramValue := m.QueryTo(paramType.Name())
			paramValues = append(paramValues, reflect.ValueOf(paramValue))
		default:
			paramValues = append(paramValues, reflect.ValueOf(nil))
		}
	}

	// 根据参数解析json等
	infoFunc.Call(paramValues)
}

func (m *MvcMiddleware) BodyToJson(obj interface{}) {
	body, _ := ioutil.ReadAll(m.Context.Request.Body)
	if err := json.Unmarshal(body, &obj); err != nil {
		panic("反序列化错误")
	}
}

func (m *MvcMiddleware) QueryTo(name string) string {
	value := m.Context.Request.URL.Query()[name][0]
	return value
}

// 根据注入的控制器构建路由
func (m *MvcMiddleware) CreateRouter() {

	routers := make([]router.ControllerDescriptor, len(m.controllerServices.controllers))
	// 控制器路由
	controllerRouter := sync.Map{}
	// 静态路由
	staticRouter := sync.Map{}

	for _, controller := range m.controllerServices.controllers {

		controllerDescriptor := router.ControllerDescriptor{
			Url:        reflect.TypeOf(controller).Name(),
			Controller: controller,
		}

		r := &router.RouterBuilder{}
		controller.Init(r)

		// 区分是控制器 Action 还是静态 Action
		nostatic, static := m.filtrate(r.RouterTable)
		controllerDescriptor.ActionRouters = nostatic
		// 写入路由树
		for _, ac := range static {
			staticRouter.Store(ac.Url, &router.StaticRouter{
				Url:        ac.Url,
				Controller: controller,
				Action:     ac,
			})
		}
	}

	for name, controller := range routers {
		controllerRouter.Store(name, &controller)
	}

	m.RouterTree = router.RouterTree{
		ControllerRouter: controllerRouter,
		StaticRouter:     staticRouter,
	}
}

// 区分静态路由还是控制器路由
func (m *MvcMiddleware) filtrate(allDescriptor map[string]router.ActionDescriptor) ([]router.ActionDescriptor, []router.ActionDescriptor) {
	controllers := make([]router.ActionDescriptor, 0, len(allDescriptor))
	actions := make([]router.ActionDescriptor, 0, len(allDescriptor))

	for _, action := range allDescriptor {
		if action.Static {
			actions = append(actions, action)
		} else {
			controllers = append(controllers, action)
		}
	}
	return controllers, actions
}
