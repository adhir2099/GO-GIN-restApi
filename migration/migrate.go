package main

import (
	"backend/configs"
	"backend/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.User{})
}
