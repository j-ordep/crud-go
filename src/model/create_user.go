package model

import (
	"fmt"

	"github.com/j-ordep/crud-go/src/config/logger"
	"github.com/j-ordep/crud-go/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("jorney", "createUser"))

	ud.EncryptPassword()

	fmt.Println()

	return nil
}
