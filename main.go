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
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *hgin.Context) {
		c.JSON(http.StatusOK, hgin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
