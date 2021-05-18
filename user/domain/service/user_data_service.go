package service

import (
	"errors"
	"github.com/hjldev/newmicro-mall/user/domain/model"
	"github.com/hjldev/newmicro-mall/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(username string, pwd string) (isOk bool, err error)
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{
		UserRepository: userRepository,
	}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func GeneratePwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func ValidatePwd(pwd string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (id int64, err error) {
	pwdByte, err := GeneratePwd(user.Pwd)
	if err != nil {
		return 0, err
	}
	user.Pwd = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(id int64) (err error) {
	return u.UserRepository.DeleteUserById(id)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		pwdByte, err := GeneratePwd(user.Pwd)
		if err != nil {
			return err
		}
		user.Pwd = string(pwdByte)
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
	return ValidatePwd(user.Pwd, pwd)
}
