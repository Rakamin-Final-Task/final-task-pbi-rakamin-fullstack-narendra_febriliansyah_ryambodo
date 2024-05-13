package main

import (
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/initializers"
	"final-task-pbi-rakamin-fullstack-narendra_febriliansyah_ryambodo/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Photo{})
}
