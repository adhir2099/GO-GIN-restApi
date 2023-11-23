package controllers

import (
	"backend/configs"
	"backend/models"

	"github.com/gin-gonic/gin"
)

type userRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserCreate(c *gin.Context) {

	body := userRequestBody{}

	c.BindJSON(&body)

	user := &models.User{Name: body.Name, Email: body.Email}

	result := configs.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to add user"})
		return
	}

	c.JSON(200, &user)
}

func UserGet(c *gin.Context) {
	var users []models.User
	configs.DB.Find(&users)
	c.JSON(200, &users)
	return
}

func UserGetById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	configs.DB.First(&user, id)
	c.JSON(200, &user)
	return
}

func UserUpdate(c *gin.Context) {

	id := c.Param("id")
	var user models.User
	configs.DB.First(&user, id)

	body := userRequestBody{}
	c.BindJSON(&body)
	data := &models.User{Name: body.Name, Email: body.Email}

	result := configs.DB.Model(&user).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update user"})
		return
	}

	c.JSON(200, &user)
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	configs.DB.Delete(&user, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
