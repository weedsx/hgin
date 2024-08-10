## hgin
一个模仿 gin 的简易 web 框架，包含 路由 (路由前缀树)、上下文、路由分组、使用中间件、模板渲染、静态资源服务、全局错误处理 (打印堆栈信息)

## 使用
> 未发布
1. 下载 hgin 包 (pub 分支) 到项目目录下，在 go.mod 引入
```
require hgin v0.0.0

replace hgin => ./hgin
```
2. 使用
```go
func main() {
    r := hgin.Default()
    r.GET("/", func(c *hgin.Context) {
        c.String(http.StatusOK, "Hello hgin")
    })
    r.GET("/panic", func(c *hgin.Context) {
        names := []string{"hgin"}
        c.String(http.StatusOK, names[100])
    })

    v2 := r.Group("/v2")
    {
        v2.GET("/hello/:name", func(c *hgin.Context) {
            c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
        })
    }

    _ = r.Run(":9999")
}
```
3. 访问 `http://localhost:9999/`

## 架构
<img src="https://raw.githubusercontent.com/weedsx/picgo/master/hgin.png"/>