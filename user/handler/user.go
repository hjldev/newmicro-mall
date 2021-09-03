package handler

import (
	"context"
	"github.com/hjldev/newmicro-mall/user/domain/model"
	"github.com/hjldev/newmicro-mall/user/domain/service"
	userPb "github.com/hjldev/newmicro-mall/user/proto/user"
	"time"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) UpdateUserInfo(ctx context.Context, request *userPb.UserInfoRequest, response *userPb.UserInfoResponse) error {
	user := &model.User{
		Id:       request.UserId,
		NickName: request.Nickname,
		Password: request.Password,
	}
	u.UserDataService.UpdateUser(user, false)
	return nil
}

func (u *User) Register(ctx context.Context, in *userPb.UserInfoRequest, response *userPb.RegisterResponse) (err error) {
	userRegister := &model.User{
		LoginName:     in.LoginName,
		NickName:      in.LoginName,
		Password:      in.Password,
		CreateTime:    time.Now(),
		IntroduceSign: "勇敢牛牛，不怕困难",
	}
	_, err = u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "成功"
	return nil
}

func (u *User) Login(ctx context.Context, in *userPb.LoginRequest, response *userPb.LoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(in.LoginName, in.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}
func (u *User) GetUserInfo(ctx context.Context, in *userPb.UserInfoRequest, out *userPb.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(in.LoginName)
	if err != nil {
		return err
	}
	out = UserForResponse(userInfo)
	return nil
}

//类型转化
func UserForResponse(userModel *model.User) *userPb.UserInfoResponse {
	response := &userPb.UserInfoResponse{}
	response.Nickname = userModel.NickName
	response.Id = userModel.Id
	return response
}
