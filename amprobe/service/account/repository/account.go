// Package repository
// Date       : 2024/9/4 14:56
// Author     : Amu
// Description:
package repository

import (
	"amprobe/pkg/database"
	"amprobe/service/model"
	"amprobe/service/schema"
	"context"
	"github.com/google/wire"
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

func NewAccountRepository(db *database.DB) IAccountRepository {
	return &AccountRepository{DB: db}
}

func (a AccountRepository) UserQuery(ctx context.Context, args schema.UserQueryArgs) (model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) UserCreate(ctx context.Context, args schema.UserCreateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) UserUpdate(ctx context.Context, args schema.UserUpdateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) UserDelete(ctx context.Context, args schema.UserDeleteArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) RoleQuery(ctx context.Context, args schema.RoleQueryArgs) (model.Roles, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) RoleCreate(ctx context.Context, args schema.RoleCreateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) RoleUpdate(ctx context.Context, args schema.RoleUpdateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) RoleDelete(ctx context.Context, args schema.RoleDeleteArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) ResourceQuery(ctx context.Context, args schema.ResourceQueryArgs) (model.Resources, error) {
	//TODO implement me
	panic("implement me")
}
