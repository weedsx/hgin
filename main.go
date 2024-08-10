package main

import (
	"hgin"
	"net/http"
)

func main() {
	StartServer()
}

func StartServer() {
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
