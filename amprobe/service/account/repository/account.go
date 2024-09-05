// Package repository
// Date       : 2024/9/4 14:56
// Author     : Amu
// Description:
package repository

import (
	"amprobe/pkg/database"
	"amprobe/pkg/utils/hash"
	"amprobe/pkg/utils/uuid"
	"amprobe/service/model"
	"amprobe/service/schema"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var AccountRepositorySet = wire.NewSet(NewAccountRepository, wire.Bind(new(IAccountRepository), new(*AccountRepository)))

type IAccountRepository interface {
	UserQuery(context.Context, schema.UserQueryArgs) (model.Users, error)
	UserCreate(context.Context, schema.UserCreateArgs) error
	UserUpdate(context.Context, schema.UserUpdateArgs) error
	UserDelete(context.Context, schema.UserDeleteArgs) error

	RoleQuery(context.Context, schema.RoleQueryArgs) (model.Roles, error)
	RoleCreate(context.Context, schema.RoleCreateArgs) error
	RoleUpdate(context.Context, schema.RoleUpdateArgs) error
	RoleDelete(context.Context, schema.RoleDeleteArgs) error

	ResourceQuery(context.Context, schema.ResourceQueryArgs) (model.Resources, error)
}

type AccountRepository struct {
	DB *database.DB
}

func NewAccountRepository(db *database.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (a *AccountRepository) UserQuery(ctx context.Context, args schema.UserQueryArgs) (model.Users, error) {
	// FIXME: 根据条件查询
	var users model.Users
	if err := a.DB.Model(&model.User{}).Preload("Roles").Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (a *AccountRepository) UserCreate(ctx context.Context, args schema.UserCreateArgs) error {
	// FIXME: 权限模型更新
	err := a.DB.RunInTransaction(func(tx *gorm.DB) error {
		u := &model.User{
			ID:       uuid.MustUUID(),
			Username: args.Username,
			Password: hash.SHA1String(args.Password),
			Status:   args.Status,
			IsAdmin:  args.IsAdmin,
			Remark:   args.Remark,
		}
		if len(args.RoleIDs) > 0 {
			var roles model.Roles
			if err := tx.Where("id IN (?)", args.RoleIDs).Find(&roles).Error; err != nil {
				return err
			}
			u.Roles = roles
		}
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a *AccountRepository) UserUpdate(ctx context.Context, args schema.UserUpdateArgs) error {
	// FIXME: 权限模型更新
	params := make(map[string]interface{})
	if args.Status != nil {
		params["status"] = *args.Status
	}
	if args.Remark != nil {
		params["remark"] = *args.Remark
	}
	if args.IsAdmin != nil {
		params["is_admin"] = *args.IsAdmin
	}
	if args.Username != nil {
		params["username"] = *args.Username
	}
	err := a.DB.RunInTransaction(func(tx *gorm.DB) error {
		if len(args.RoleIDs) > 0 {
			var roles model.Roles
			if err := tx.Where("id IN (?)", args.RoleIDs).Find(&roles).Error; err != nil {
				return err
			}
			params["roles"] = roles
		}
		err := tx.Model(&model.User{}).Association("Roles").Clear()
		if err != nil {
			return err
		}
		if err := tx.Model(&model.User{}).Where("id = ?", args.ID).Updates(params).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a *AccountRepository) UserDelete(ctx context.Context, args schema.UserDeleteArgs) error {
	// FIXME: 权限模型更新
	err := a.DB.RunInTransaction(func(tx *gorm.DB) error {
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountRepository) RoleQuery(ctx context.Context, args schema.RoleQueryArgs) (model.Roles, error) {
	// FIXME: 权限模型更新
	var roles model.Roles
	if err := a.DB.Model(&model.Role{}).Preload("Resources").Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}

func (a *AccountRepository) RoleCreate(ctx context.Context, args schema.RoleCreateArgs) error {
	return a.DB.RunInTransaction(func(tx *gorm.DB) error {
		return nil
	})
}

func (a *AccountRepository) RoleUpdate(ctx context.Context, args schema.RoleUpdateArgs) error {
	// FIXME: 权限模型更新
	return a.DB.RunInTransaction(func(tx *gorm.DB) error {
		return nil
	})
}

func (a *AccountRepository) RoleDelete(ctx context.Context, args schema.RoleDeleteArgs) error {
	// FIXME: 权限模型更新
	return a.DB.RunInTransaction(func(tx *gorm.DB) error {
		return nil
	})
}

func (a *AccountRepository) ResourceQuery(ctx context.Context, args schema.ResourceQueryArgs) (model.Resources, error) {
	var resources model.Resources
	if err := a.DB.Model(&model.Resource{}).Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&resources).Error; err != nil {
		return resources, err
	}
	return resources, nil
}
