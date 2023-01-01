package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	app.SetTrustedProxies([]string{"127.0.0.1"})
}
