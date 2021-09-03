package handler

import (
	"context"
	"github.com/hjldev/newmicro-mall/user/proto/user"
	"github.com/prometheus/common/log"
)

type UserApi struct {
	UserService user.UserService
}

func (u UserApi) Register(ctx context.Context, request *user.UserInfoRequest, response *user.RegisterResponse) error {
	log.Info("接受到 访问请求")
	return nil
}

func (u UserApi) Login(ctx context.Context, request *user.LoginRequest, response *user.LoginResponse) error {
	panic("implement me")
}

func (u UserApi) GetUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	panic("implement me")
}

func (u UserApi) UpdateUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	panic("implement me")
}
