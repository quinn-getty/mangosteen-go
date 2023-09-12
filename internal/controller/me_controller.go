package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
func (ctrl *MeController) Get(ctx *gin.Context) {
	auth := ctx.GetHeader(Authorization)
	if len(auth) < 8 {
		log.Print("无效的JWT Authorization < 8", auth)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	jwtString := auth[7:]
	token, err := jwt_helper.Parse(jwtString)
	if err != nil {
		log.Print("无效的JWT jwtString jwt_helper.Parse 失败", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	m, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Print("无效的JWT token.Claims 失败", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	userIdF64, ok := m["user_id"].(float64)
	userId := int32(userIdF64)
	if !ok {
		log.Print("无效的JWT userIdStr 获取失败", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	user, err := q.FindUserById(database.DBCtx, userId)
	if err != nil {
		log.Println(userId)
		log.Print("无效的JWT FindUserById ", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	// ctx.JSON(http.StatusOK, GetMeResBody{
	// 	Resourse: user,
	// })

	ctx.JSON(http.StatusOK, gin.H{
		"resourse": user,
	})

}

func (ctrl *MeController) Create(c *gin.Context)  {}
func (ctrl *MeController) Delete(c *gin.Context)  {}
func (ctrl *MeController) Update(c *gin.Context)  {}
func (ctrl *MeController) getList(c *gin.Context) {}
