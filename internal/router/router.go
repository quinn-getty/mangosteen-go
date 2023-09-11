package router

import (
	"mangosteen/config"
	"mangosteen/docs"
	"mangosteen/internal/controller"
	"mangosteen/internal/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth(JWT)

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func New() *gin.Engine {
	config.LoadConfig()

	r := gin.Default()
	docs.SwaggerInfo.Version = "1.0"
	database.Connect()

	api := r.Group("/api")
	apiV1 := api.Group("/v1")

	session := controller.SessionController{}
	session.RegisterRouter(apiV1)

	aildationCode := controller.VaildationCodeController{}
	aildationCode.RegisterRouter(apiV1)

	r.GET("/api/v1/ping", controller.Ping)
	// r.POST("/api/v1/create_validation_code", controller.CreateValidationCode)
	// r.POST("/api/v1/session", controller.CreateSession)

	// swager router
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
