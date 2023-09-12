package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRouter(rg *gin.RouterGroup)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	getList(c *gin.Context)
}

var Authorization = "Authorization"
