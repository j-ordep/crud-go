package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j-ordep/crud-go/src/config/logger"
	"github.com/j-ordep/crud-go/src/config/validation"
	"github.com/j-ordep/crud-go/src/controller/model/request"
	"github.com/j-ordep/crud-go/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("jorney","createUser"),
	)

	var userRequest request.UserRequest

	// ShouldBindJSON já faz validação (validate:"required")
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}
	
	logger.Info("user created", zap.String("jorney","createUser"))

	c.JSON(http.StatusOK, response)
}