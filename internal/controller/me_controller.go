package controller

import (
	"log"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MeController struct {
}

func (ctrl *MeController) RegisterRouter(rg *gin.RouterGroup) {
	me := rg.Group("/me")
	me.GET(",", ctrl.Get)
}

func (ctrl *MeController) Get(ctx *gin.Context) {
	auth := ctx.GetHeader(Authorization)
	jwtString := auth[7:]
	token, err := jwt_helper.Parse(jwtString)
	if err != nil {
		log.Print("无效的JWT", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	m, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Print("无效的JWT", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	userIdStr, ok := m["user_id"].(string)

	if !ok {
		log.Print("无效的JWT", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	userIdInt64, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Print("无效的JWT", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	userId := int32(userIdInt64)
	q := database.NewQuery()

	user, err := q.FindUserById(database.DBCtx, userId)
	if err != nil {
		log.Print("无效的JWT", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"resourse": user,
	})

}

func (ctrl *MeController) Create(c *gin.Context)  {}
func (ctrl *MeController) Delete(c *gin.Context)  {}
func (ctrl *MeController) Update(c *gin.Context)  {}
func (ctrl *MeController) getList(c *gin.Context) {}
