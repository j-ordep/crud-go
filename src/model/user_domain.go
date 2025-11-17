package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/j-ordep/crud-go/src/config/rest_err"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email: email,
		Password: password,
		Name: name,
		Age: age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

//aula 10, 12min