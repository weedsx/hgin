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
		c.String(http.StatusOK, "Hello Howard\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *hgin.Context) {
		names := []string{"Howard"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8080")
}
