// Package middleware
// Date: 2024/3/27 16:29
// Author: Amu
// Description:
package middleware

import (
	"runtime"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

var defaultStackTraceBufLen = 2048

func StackTraceHandler(c *fiber.Ctx, e interface{}) {
	buf := make([]byte, defaultStackTraceBufLen)
	buf = buf[:runtime.Stack(buf, true)]
	_, _ = c.Write(buf)
}

var StackMiddleware = recover.New(recover.Config{
	EnableStackTrace:  true,
	StackTraceHandler: StackTraceHandler,
})
