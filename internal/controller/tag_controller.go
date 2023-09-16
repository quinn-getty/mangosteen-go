package controller

import "github.com/gin-gonic/gin"

type TagController struct {
}

func (ctrl *TagController) RegisterRouter(rg *gin.RouterGroup) {
	item := rg.Group("/tag")
	item.POST("", ctrl.Create)
	item.GET("", ctrl.getList)

}
func (ctrl *TagController) Create(c *gin.Context)  {}
func (ctrl *TagController) Delete(c *gin.Context)  {}
func (ctrl *TagController) Update(c *gin.Context)  {}
func (ctrl *TagController) Get(c *gin.Context)     {}
func (ctrl *TagController) getList(c *gin.Context) {}
