package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
}

type CreateSessionReqBody struct {
	Email string `json:"email" binging:"required"`
	Code  string `json:"code" binging:"required"`
}

type CreateSessionResBody struct {
	JWT    string `json:"jwt"`
	UserId int32  `json:"userId"`
}

func (ctrl *SessionController) RegisterRouter(rg *gin.RouterGroup) {
	session := rg.Group("/session")
	session.POST("", ctrl.Create)
}

// 登录 godoc
// @Summary      session
// @Description  获取session
// @Tags         session
// @Accept       json
// @Produce      json
// @Params request body CreateSessionReqBody true "query params"
// @Success      200 {object} CreateSessionResBody
// @Failure      500
// @Router       /session [post]
func (ctrl *SessionController) Create(ctx *gin.Context) {
	var reqBody CreateSessionReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		log.Println("入参错误")
		ctx.String(http.StatusBadRequest, "参数错误")
		return
	}

	q := database.NewQuery()
	_, err := q.FindValidationCode(database.DBCtx, queries.FindValidationCodeParams{
		Email: reqBody.Email,
		Code:  reqBody.Code,
	})
	if err != nil {
		log.Println("没有查询到验证码")
		ctx.String(http.StatusBadRequest, "验证码错误")
	}

	user, err := q.FindUserByEmail(database.DBCtx, reqBody.Email)
	if err != nil {
		user, err = q.CreateUser(database.DBCtx, reqBody.Email)
		if err != nil {
			log.Println("创建失败")
			ctx.String(http.StatusInternalServerError, "稍后重试")
		}
	}

	jwt, err := jwt_helper.GenerateJWT(1)
	if err != nil {
		log.Println("生成jwt失败")
		ctx.String(http.StatusInternalServerError, "稍后重试")
	}

	resBody := CreateSessionResBody{
		JWT:    jwt,
		UserId: user.ID,
	}

	ctx.JSON(http.StatusOK, resBody)
}

func (ctrl *SessionController) Delete(c *gin.Context)  {}
func (ctrl *SessionController) Update(c *gin.Context)  {}
func (ctrl *SessionController) getOne(c *gin.Context)  {}
func (ctrl *SessionController) getList(c *gin.Context) {}
