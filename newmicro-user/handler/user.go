package handler

import (
	"context"
	"github.com/hjldev/newmicro-mall/newmicro-user/domain/model"
	"github.com/hjldev/newmicro-mall/newmicro-user/domain/service"
	userPb "github.com/hjldev/newmicro-mall/newmicro-user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, in *userPb.RegisterRequest, response *userPb.RegisterResponse) (err error) {
	userRegister := &model.User{
		UserName:  in.UserName,
		FirstName: in.FirstName,
		Pwd:       in.Pwd,
	}
	_, err = u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "成功"
	return nil
}

func (u *User) Login(ctx context.Context, in *userPb.LoginRequest, response *userPb.LoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(in.UserName, in.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}
func (u *User) GetUserInfo(ctx context.Context, in *userPb.UserInfoRequest, out *userPb.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(in.UserName)
	if err != nil {
		return err
	}
	out = UserForResponse(userInfo)
	return nil
}

//类型转化
func UserForResponse(userModel *model.User) *userPb.UserInfoResponse {
	response := &userPb.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.Id = userModel.Id
	return response
}
