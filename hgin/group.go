package hgin

import (
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share an Engine instance
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	nextGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, nextGroup)
	return nextGroup
}

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Static 方法用于注册静态文件路由。
// 该方法允许通过相对路径和根目录提供静态文件服务。
// 参数 relativePath 是静态资源的访问路径，root 是静态文件的根目录。
// 通过这个方法，可以方便地为特定路径下的所有静态文件提供服务。
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}

// createStaticHandler 创建一个静态文件处理的路由规则
// 该函数属于RouterGroup结构体的方法，用于在路由组中添加静态文件处理功能。
// relativePath: 相对于路由组前缀的路径；
// fs: http.FileSystem接口，用于访问文件系统；
// HandlerFunc: 一个处理HTTP请求的函数。
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	// 假设 absolutePath 是 /static，那么如果客户端请求 /static/js/app.js，
	// http.StripPrefix("/static", ...) 会去掉 /static 部分，只剩下 /js/app.js，
	// 然后将其传递给 http.FileServer(fs) 进行处理。
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Params["filepath"]
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}
