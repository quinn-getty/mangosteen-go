package router

import (
	"mangosteen/docs"
	"mangosteen/internal"
	"mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadController(rg *gin.RouterGroup) {
	controllerList := []controller.Controller{
		&controller.SessionController{},
		&controller.VaildationCodeController{},
		&controller.MeController{},
		&controller.ItemController{},
		&controller.TagController{},
	}

	for _, c := range controllerList {
		c.RegisterRouter(rg)
	}

}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func New() *gin.Engine {
	r := gin.Default()
	internal.InitRouter(r)

	docs.SwaggerInfo.Version = "1.0"
	api := r.Group("/api")
	apiV1 := api.Group("/v1")

	LoadController(apiV1)

	r.GET("/api/v1/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
