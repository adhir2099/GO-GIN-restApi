package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRouter(router *gin.Engine) {

	routes := router.Group("api/v1/users")
	routes.POST("", controllers.UserCreate)
	routes.GET("", controllers.UserGet)
	routes.GET("/:id", controllers.UserGetById)
	routes.PUT("/:id", controllers.UserUpdate)
	routes.DELETE("/:id", controllers.UserDelete)

}
