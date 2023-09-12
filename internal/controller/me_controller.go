package controller

import (
	"mangosteen/config/queries"
	"mangosteen/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MeController struct {
}

type GetMeResBody struct {
	Resourse queries.User `json:"resourse"`
}

func (ctrl *MeController) RegisterRouter(rg *gin.RouterGroup) {
	me := rg.Group("/me")
	me.GET("", ctrl.Get)
}

// GetMe godoc
// @Summary      获取当前用户
// @Description  获取当前用户信息
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200 {object} GetMeResBody
// @Failure      500
// @Router       /me [get]
func (ctrl *MeController) Get(c *gin.Context) {
	user, ok := middleware.GetMe(c)

	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	resBody := GetMeResBody{
		Resourse: user,
	}

	c.JSON(http.StatusOK, resBody)

}

func (ctrl *MeController) Create(c *gin.Context)  {}
func (ctrl *MeController) Delete(c *gin.Context)  {}
func (ctrl *MeController) Update(c *gin.Context)  {}
func (ctrl *MeController) getList(c *gin.Context) {}
