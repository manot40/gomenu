package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/manot40/gomenu/internal/middlewares"
	"github.com/manot40/gomenu/internal/models"
	"github.com/manot40/gomenu/internal/routes"
)

func Serve() {
	app := gin.Default()

	models.ConnectDatabase()

	middlewares.Init(app)

	routes.Init(app)

	app.Run()
}
