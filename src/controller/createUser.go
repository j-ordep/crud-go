package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j-ordep/crud-go/src/config/rest_err"
	"github.com/j-ordep/crud-go/src/model"
)

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Erro ao bindar usuario"))
	}

	err := rest_err.NewBadRequestError("Voce chamou a rota de forma errada")

	c.JSON(err.Code, err)
}