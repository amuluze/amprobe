// Package service
// Date: 2024/3/27 16:39
// Author: Amu
// Description:
package service

import (
	"context"
	"github.com/amuluze/amprobe/pkg/auth"
	"github.com/amuluze/amprobe/service/auth/repository"
	"github.com/amuluze/amprobe/service/schema"
	"github.com/amuluze/amutool/errors"
	"github.com/google/wire"
	"log/slog"
)

var AuthServiceSet = wire.NewSet(NewAuthService, wire.Bind(new(IAuthService), new(*AuthService)))

type IAuthService interface {
	Login(ctx context.Context, args *schema.LoginArgs) (*schema.LoginResult, error)
	Logout(ctx context.Context, userID, token string) error
	PassUpdate(ctx context.Context, args *schema.PasswordUpdateArgs) error
	TokenUpdate(ctx context.Context, token string) (*schema.LoginResult, error)
}

type AuthService struct {
	Auth     auth.Auther
	AuthRepo *repository.AuthRepo
}

func NewAuthService(auth auth.Auther, authRepo *repository.AuthRepo) *AuthService {
	return &AuthService{Auth: auth, AuthRepo: authRepo}
}

func (a *AuthService) Login(ctx context.Context, args *schema.LoginArgs) (*schema.LoginResult, error) {
	u, err := a.AuthRepo.Login(ctx, args)
	if err != nil {
		slog.Error("auth repo login failed", "error", err)
		return nil, errors.New400Error(err.Error())
	}
	tokenInfo, err := a.Auth.GenerateToken(u.ID.String(), u.Username, u.IsAdmin)
	if err != nil {
		slog.Error("generate token failed", "error", err)
		return nil, errors.New400Error(err.Error())
	}
	res := &schema.LoginResult{
		AccessToken:  tokenInfo.GetAccessToken(),
		RefreshToken: tokenInfo.GetRefreshToken(),
	}
	return res, nil
}

func (a *AuthService) Logout(ctx context.Context, userID, token string) error {
	if userID != "" {
		err := a.Auth.DestroyToken(token)
		if err != nil {
			slog.Error("destroy token failed", "error", err)
			return errors.New400Error(err.Error())
		}
	}
	return nil
}

func (a *AuthService) PassUpdate(ctx context.Context, args *schema.PasswordUpdateArgs) error {
	if args.OldPassword == args.NewPassword {
		slog.Error("old password equal new password")
		return errors.New400Error("equal password")
	}
	err := a.AuthRepo.PassUpdate(ctx, args)
	if err != nil {
		slog.Error("auth pass update failed", "error", err)
		return errors.New400Error(err.Error())
	}
	return nil
}

func (a *AuthService) TokenUpdate(ctx context.Context, token string) (*schema.LoginResult, error) {
	userID, username, isAdmin, err := a.Auth.ParseToken(token, "refresh_token")
	if err != nil {
		slog.Error("parse token failed", "error", err)
		return nil, errors.New400Error(err.Error())
	}
	slog.Info("token update", "user_id", userID)
	tokenInfo, err := a.Auth.GenerateToken(userID, username, isAdmin)
	if err != nil {
		slog.Error("generate new token failed", "error", err)
		return nil, errors.New400Error(err.Error())
	}

	res := &schema.LoginResult{
		AccessToken:  tokenInfo.GetAccessToken(),
		RefreshToken: tokenInfo.GetRefreshToken(),
	}

	return res, nil
}
