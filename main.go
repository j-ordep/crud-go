package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/j-ordep/crud-go/src/controller/routes"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default() // inicia o gin com um logger e middleware

	routes.InitRoutes(&router.RouterGroup)	

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}