// Package fiberx
// Date: 2024/3/6 13:15
// Author: Amu
// Description:
package fiberx

import (
	"log/slog"
	"net/http"
	"strings"
	
	"amprobe/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetToken Get jwt token from header (Authorization: Bearer xxx)
func GetToken(c *fiber.Ctx) string {
	var token string
	auth := c.GetReqHeaders()["Authorization"][0]
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// Success response.status = 200
func Success(c *fiber.Ctx, v interface{}) error {
	return ReturnJson(c, http.StatusOK, v)
}

type FailedResponse struct {
	Err string `json:"err"` // 响应错误，来自 service 层的错误信息
	Msg string `json:"msg"` // 错误消息，来自 api 层的错误信息
}

// Failure response.status = 400
func Failure(c *fiber.Ctx, err errors.Error) error {
	// 正常来说api层返回给接口层的错误信息不用打印，在这里会进行统一打印，400-500之间会打印warn 日志
	// 500 会打印error 日志
	if status := err.Status; status >= 400 && status < 500 {
		slog.Warn("return 4xx error", "error", err)
	} else if status >= 500 {
		slog.Error("return 5xx error", "error", err)
	}
	
	return ReturnJson(c, err.Status, &FailedResponse{Err: err.Err, Msg: err.Msg})
}

// Unauthorized response.status = 401 when token error or token is fail
func Unauthorized(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusUnauthorized)
}

// NoContent response.status = 204
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNoContent)
}

// Forbidden response.status = 403 when permission error
func Forbidden(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusForbidden)
}

func ReturnJson(c *fiber.Ctx, status int, v interface{}) error {
	c.Status(status)
	return c.JSON(v)
}

// ParseQuery Parse query parameter to struct
func ParseQuery(c *fiber.Ctx, obj interface{}) error {
	if err := c.QueryParser(obj); err != nil {
		return err
	}
	return nil
}

func ParseBody(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return err
	}
	return nil
}
