package main

import (
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/controllers"
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/users/register", controllers.UserRegisterController)
	r.GET("/users/login", controllers.UserLoginController)
	r.PUT("/users/update", controllers.UserUpdateController)
	r.DELETE("/users/delete", controllers.UserDeleteController)
	r.Run()
}
