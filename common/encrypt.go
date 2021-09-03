package common

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//加密
func Encrypt(pwd string) (encryptPwd string, err error) {
	if pwd == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		pwd = string(hash)
		return pwd, err
	}
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
