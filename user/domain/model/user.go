package model

import "time"

type User struct {
	CreateTime    time.Time `gorm:"column:create_time;not null" json:"createTime"`
	IntroduceSign string    `gorm:"column:introduce_sign;not null" json:"introduceSign"`
	IsDeleted     int       `gorm:"column:is_deleted;not null" json:"isDeleted"`
	LockedFlag    int       `gorm:"column:locked_flag;not null" json:"lockedFlag"`
	LoginName     string    `gorm:"column:login_name;not null" json:"loginName"`
	NickName      string    `gorm:"column:nick_name;not null" json:"nickName"`
	PasswordMd5   string    `gorm:"column:password_md5;not null" json:"passwordMd5"`
	Id            int64     `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	Password      string    `gorm:"-" json:"-"`
}

func (model *User) TableName() string {
	return "tb_user"
}
