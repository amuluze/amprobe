// Package middleware
// Date       : 2024/8/27 18:30
// Author     : Amu
// Description:
package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if SkipHandler(c, skippers...) {
			return c.Next()
		}
		//ctx := c.UserContext()
		//
		//path := c.Path()
		//method := c.Method()

		return c.Next()
	}
}
