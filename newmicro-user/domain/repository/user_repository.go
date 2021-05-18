package repository

import (
	"github.com/hjldev/newmicro-mall/newmicro-user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindUserByName(string) (*model.User, error)
	FindUserId(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserById(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		mysqlDb: db,
	}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

func (u *UserRepository) FindUserId(id int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, id).Error
}

func (u *UserRepository) CreateUser(user *model.User) (id int64, err error) {
	return user.Id, u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserById(id int64) error {
	return u.mysqlDb.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Save(user).Error
}

func (u *UserRepository) FindAll() (users []model.User, err error) {
	return users, u.mysqlDb.Find(&users).Error
}
