// Package service
// Date: 2024/3/27 17:04
// Author: Amu
// Description:
package service

import (
	"amprobe/pkg/database"
	"amprobe/pkg/utils/hash"
	"amprobe/service/model"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log/slog"
	
	"github.com/google/wire"
)

var PrepareSet = wire.NewSet(wire.Struct(new(Prepare), "*"))

var users = []*model.User{
	{
		Username: "admin",
		Password: "admin123",
		Remark:   "",
		IsAdmin:  1,
		Status:   1,
		Roles: []*model.Role{
			{
				Name:   "管理员",
				Status: 1,
			},
		},
	},
	{
		Username: "amprobe",
		Password: "123456",
		Remark:   "",
		IsAdmin:  1,
		Status:   1,
		Roles: []*model.Role{
			{
				Name:   "普通用户",
				Status: 1,
			},
		},
	},
}

type Prepare struct {
	db       *database.DB
	enforcer *casbin.SyncedEnforcer
}

type NamePolicy struct {
	RoleID string
	Path   string
	Method string
}

type GroupPolicy struct {
	UserID string
	RoleID string
}

func (a *Prepare) Init(app *fiber.App) {
	// init account
	a.InitAccount(app)
	
	// init casbin rules
	a.InitCasbinRules()
}

func (a *Prepare) InitAccount(app *fiber.App) {
	var getResources []*model.Resource
	var postResources []*model.Resource
	
	for _, routers := range app.Stack() {
		for _, router := range routers {
			postResources = append(postResources, &model.Resource{
				Name:   router.Name,
				Path:   router.Path,
				Method: router.Method,
				Status: 1,
			})
			if router.Method == "GET" || router.Name == "登录" || router.Name == "登出" || router.Name == "更新密码" || router.Name == "更新 token" {
				getResources = append(getResources, &model.Resource{
					Name:   router.Name,
					Path:   router.Path,
					Method: router.Method,
					Status: 1,
				})
			}
		}
	}
	
	slog.Info("get resources", "resources", getResources)
	slog.Info("post resources", "resources", postResources)
	
	_ = a.db.RunInTransaction(func(tx *gorm.DB) error {
		for _, u := range users {
			for _, role := range u.Roles {
				if role.Name == "管理员" {
					role.Resources = postResources
				} else {
					role.Resources = getResources
				}
			}
			var ou model.User
			// 不存在则创建
			if err := a.db.Model(&model.User{}).Where("username = ?", u.Username).Take(&ou).Error; err != nil {
				a.db.Create(u)
			} else {
				// 存在则更新
				a.db.Model(&model.User{}).Where("username = ?", u.Username).Updates(model.User{Password: hash.SHA1String(u.Password), Status: u.Status, IsAdmin: u.IsAdmin})
			}
		}
		return nil
	})
}

func (a *Prepare) InitCasbinRules() {
	var users []*model.User
	if err := a.db.Preload("Role").Preload("Resource").Find(&users).Error; err != nil {
		slog.Error("get all users error", "error", err)
		return
	}
	
	for _, user := range users {
		for _, role := range user.Roles {
			if _, err := a.enforcer.AddNamedGroupingPolicy("g", user.ID, role.ID); err != nil {
				slog.Error("add grouping policy error", "error", err)
			}
			for _, resource := range role.Resources {
				if _, err := a.enforcer.AddNamedPolicy("p", role.ID, resource.Path, resource.Method); err != nil {
					slog.Error("add policy error", "error", err)
				}
			}
		}
	}
}
