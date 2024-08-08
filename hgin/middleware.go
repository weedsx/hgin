package hgin

import (
	"log"
	"time"
)

// Logger 记录请求日志
func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s - %s in %v ms", c.StatusCode, c.Method, c.Req.RequestURI, time.Since(t).Milliseconds())
	}
}
