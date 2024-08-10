package hgin

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
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

// Recovery 处理 panic 并打印堆栈信息
func Recovery() HandlerFunc {
	return func(c *Context) {
		// defer 处理 panic 并打印堆栈信息
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(msg))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

// trace 打印堆栈信息
func trace(msg string) string {
	var pcs = make([]uintptr, 32)
	// Callers 用来返回调用栈的程序计数器,
	// 第 0 个 Caller 是 Callers 本身,
	// 第 1 个是上一层 trac,
	// 第 2 个是再上一层的 defer func
	// 因此，为了日志简洁一点，我们跳过了前 3 个 Caller
	n := runtime.Callers(3, pcs[:])

	var sb strings.Builder
	sb.WriteString(msg + "\nTraceback:")
	// 通过 runtime.FuncForPC(pc) 获取对应的函数,
	// 通过 fn.FileLine(pc) 获取到调用该函数的文件名和行号，打印在日志中
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return sb.String()
}
