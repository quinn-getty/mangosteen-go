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

func loadController(rg *gin.RouterGroup) {
	controllerList := []controller.Controller{
		&controller.SessionController{},
		&controller.VaildationCodeController{},
	}

	for _, c := range controllerList {
		c.RegisterRouter(rg)
	}

}

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
	loadController(apiV1)

	r.GET("/api/v1/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
