package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/j-ordep/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/user", controller.CreateUser)
	// r.GET("/user/:userId", controller.FindUserById)
	// r.GET("/user/:userEmail", controller.FindUserByEmail)
	// r.PUT("/user", controller.UpdateUser)
	// r.DELETE("/user", controller.DeleteUser)
}