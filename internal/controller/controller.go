package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRouter(rg *gin.RouterGroup)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	getOne(c *gin.Context)
	getList(c *gin.Context)
}
