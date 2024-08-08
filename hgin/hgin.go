package hgin

import (
	"net/http"
)

// HandlerFunc 定义请求处理器
type HandlerFunc func(c *Context)

// Engine 实现 ServeHTTP 接口
type Engine struct {
	router *router

	*RouterGroup // 继承 RouterGroup
	groups       []*RouterGroup
}

// New Engine 的构造器
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := newContext(w, req)
	engine.router.handle(context)
}

func (engine *Engine) addRoute(reqMethod, pattern string, handler HandlerFunc) {
	engine.router.addRoute(reqMethod, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 开启一个 HTTP 服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
