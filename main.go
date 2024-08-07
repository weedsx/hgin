package main

import (
	"hgin"
	"net/http"
)

func main() {
	r := hgin.New()
	r.GET("/", func(c *hgin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello hgin</h1>")
	})

	r.GET("/hello", func(c *hgin.Context) {
		// expect /hello?name=howard
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *hgin.Context) {
		// expect /hello/howard
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *hgin.Context) {
		c.JSON(http.StatusOK, hgin.H{"filepath": c.Param("filepath")})
	})

	_ = r.Run(":9999")
}
