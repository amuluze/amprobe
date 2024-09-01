// Package middleware
// Date: 2024/3/27 16:29
// Author: Amu
// Description:
package middleware

import (
	"encoding/json"
	"log/slog"
	"runtime"

	"github.com/amuluze/amutool/errors"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

var defaultStackTraceBufLen = 4096

func StackTraceHandler(c *fiber.Ctx, e interface{}) {
	buf := make([]byte, defaultStackTraceBufLen)
	slog.Error("e", "err", e)
	buf = buf[:runtime.Stack(buf, false)]
	slog.Error("recorver stack trace", "buf", string(buf), "err", e)
	data, _ := json.Marshal(errors.ErrInternalServer)
	_, _ = c.Write(data)
}

var StackMiddleware = recover.New(recover.Config{
	EnableStackTrace:  true,
	StackTraceHandler: StackTraceHandler,
})
