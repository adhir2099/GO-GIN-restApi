package main

import (
	"net/http"

	"backend/configs"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Running Server Go.",
		})
	})
	r.Run()
}
