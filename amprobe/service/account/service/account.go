// Package service
// Date       : 2024/9/4 14:56
// Author     : Amu
// Description:
package service

import (
	"amprobe/service/account/repository"
	"amprobe/service/schema"
	"context"
	"github.com/google/wire"
)

var AccountServiceSet = wire.NewSet(NewAccountService, wire.Bind(new(IAccountService), new(*AccountService)))

type IAccountService interface {
	UserQuery(context.Context, schema.UserQueryArgs) (schema.UserQueryReply, error)
	UserCreate(context.Context, schema.UserCreateArgs) error
	UserUpdate(context.Context, schema.UserUpdateArgs) error
	UserDelete(context.Context, schema.UserDeleteArgs) error

	RoleQuery(context.Context, schema.RoleQueryArgs) (schema.RoleQueryReply, error)
	RoleCreate(context.Context, schema.RoleCreateArgs) error
	RoleUpdate(context.Context, schema.RoleUpdateArgs) error
	RoleDelete(context.Context, schema.RoleDeleteArgs) error

	ResourceQuery(context.Context, schema.ResourceQueryArgs) (schema.ResourceQueryReply, error)
}

type AccountService struct {
	AccountRepository repository.IAccountRepository
}

func NewAccountService(accountRepository repository.IAccountRepository) *AccountService {
	return &AccountService{AccountRepository: accountRepository}
}

func (a AccountService) UserQuery(ctx context.Context, args schema.UserQueryArgs) (schema.UserQueryReply, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) UserCreate(ctx context.Context, args schema.UserCreateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) UserUpdate(ctx context.Context, args schema.UserUpdateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) UserDelete(ctx context.Context, args schema.UserDeleteArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) RoleQuery(ctx context.Context, args schema.RoleQueryArgs) (schema.RoleQueryReply, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) RoleCreate(ctx context.Context, args schema.RoleCreateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) RoleUpdate(ctx context.Context, args schema.RoleUpdateArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) RoleDelete(ctx context.Context, args schema.RoleDeleteArgs) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) ResourceQuery(ctx context.Context, args schema.ResourceQueryArgs) (schema.ResourceQueryReply, error) {
	//TODO implement me
	panic("implement me")
}
