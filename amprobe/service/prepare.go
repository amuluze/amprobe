// Package service
// Date: 2024/3/27 17:04
// Author: Amu
// Description:
package service

import (
	"amprobe/pkg/database"
	"amprobe/pkg/utils/hash"
	"amprobe/pkg/utils/uuid"
	"amprobe/service/model"
	
	"github.com/google/wire"
	"gorm.io/gorm"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Remark   string `yaml:"remark"`
	IsAdmin  int    `yaml:"is_admin"`
	Status   int    `yaml:"status"`
}

var users = []User{
	{
		Username: "admin",
		Password: "admin123",
		Remark:   "管理员",
		IsAdmin:  1,
		Status:   1,
	},
	{
		Username: "amprobe",
		Password: "123456",
		Remark:   "普通用户",
		IsAdmin:  1,
		Status:   1,
	},
}

var PrepareSet = wire.NewSet(wire.Struct(new(Prepare), "*"))

type Prepare struct {
	db *database.DB
}

func (a *Prepare) Init(config *Config) {
	_ = a.db.RunInTransaction(func(tx *gorm.DB) error {
		for _, u := range users {
			var ou model.User
			// 不存在则创建
			if err := a.db.Model(&model.User{}).Where("username = ?", u.Username).Take(&ou).Error; err != nil {
				a.db.Create(&model.User{
					ID:       uuid.MustUUID(),
					Username: u.Username,
					Password: hash.SHA1String(u.Password),
					Status:   u.Status,
					IsAdmin:  u.IsAdmin,
				})
			}
			// 存在则更新
			a.db.Model(&model.User{}).Where("username = ?", u.Username).Updates(model.User{Password: hash.SHA1String(u.Password), Status: u.Status, IsAdmin: u.IsAdmin})
		}
		return nil
	})
}
