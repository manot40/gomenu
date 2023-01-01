package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manot40/gomenu/internal/controllers"
)

func Tag(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllTag)
	router.POST("/", controllers.CreateTag)
	router.DELETE("/:id", controllers.DeleteTag)
}
