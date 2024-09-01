// Package middleware
// Date: 2024/3/27 16:20
// Author: Amu
// Description:
package middleware

import (
	"log/slog"

	"amprobe/pkg/auth"

	"amprobe/pkg/contextx"
	"amprobe/pkg/fiberx"
	"amprobe/service/schema"

	"github.com/gofiber/fiber/v2"
)

func wrapUserAuthContext(c *fiber.Ctx, userID string, username string) {
	ctx := contextx.NewUserID(c.UserContext(), userID)
	ctx = contextx.NewUsername(ctx, username)
	c.SetUserContext(ctx)
}

func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if SkipHandler(c, skippers...) {
			var args schema.LoginArgs
			_ = c.BodyParser(&args)
			if c.Path() == "/v1/auth/login" {
				a.RecordAudit(args.Username, "登录")
			}
			return c.Next()
		}

		slog.Debug("auth middleware", "token", fiberx.GetToken(c))
		var userID string
		var username string
		var isAdmin string
		var err error
		userID, username, isAdmin, err = a.ParseToken(fiberx.GetToken(c), "access_token")
		if err != nil {
			slog.Error("parse token failed", "error", err)
			return fiberx.Unauthorized(c)
		}

		wrapUserAuthContext(c, userID, username)
		// FIXME: 这里仅是一个简单的权限控制方案，禁止非管理员用户放完 post 方法
		if c.Method() == "POST" && isAdmin != "1" && c.Path() != "/v1/auth/logout" {
			return fiberx.Forbidden(c)
		}

		err = c.Next()
		if err == nil {
			if (c.Method() == "POST" && isAdmin == "1") || c.Path() == "/v1/auth/logout" {
				a.RecordAudit(username, OperateEvent[c.Path()])
			}
		}
		return err
	}
}
