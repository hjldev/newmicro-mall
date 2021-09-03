package service

import (
	"github.com/hjldev/newmicro-mall/common"
	"github.com/hjldev/newmicro-mall/user/domain/model"
	"github.com/hjldev/newmicro-mall/user/domain/repository"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(username string, pwd string) (bool, error)
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{
		UserRepository: userRepository,
	}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func (u *UserDataService) AddUser(user *model.User) (id int64, err error) {
	pwdByte := common.Md5(user.Password)
	user.PasswordMd5 = pwdByte
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(id int64) (err error) {
	return u.UserRepository.DeleteUserById(id)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		pwdByte := common.Md5(user.Password)
		user.PasswordMd5 = pwdByte
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(name string) (*model.User, error) {
	return u.UserRepository.FindUserByName(name)
}

func (u *UserDataService) CheckPwd(username string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(username)
	if err != nil {
		return false, err
	}
	userPwd := common.Md5(pwd)
	return userPwd == user.PasswordMd5, err
}
