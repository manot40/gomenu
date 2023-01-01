package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manot40/gomenu/internal/controllers"
)

func Menu(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllMenu)
	router.GET("/:id", controllers.GetMenu)
	router.POST("/", controllers.CreateMenu)
	router.PUT("/:id", controllers.UpdateMenu)
	router.DELETE("/:id", controllers.DeleteMenu)
}
