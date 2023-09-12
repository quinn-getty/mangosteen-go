package internal

import (
	"mangosteen/config"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	config.LoadConfig()
	database.Connect()
	r.Use(middleware.Me([]string{
		"/api/v1/ping",
		"/api/v1/session",
		"/api/v1/create_validation_code",
	}))
}
