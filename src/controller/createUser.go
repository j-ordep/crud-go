package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/j-ordep/crud-go/src/config/validation"
	"github.com/j-ordep/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	// ShouldBindJSON já faz validação (validate:"required")
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}