package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSessionReqBody struct {
	Email string `json:"email" binging:"required"`
	Code  string `json:"code" binging:"required"`
}

type CreateSessionResBody struct {
	JWT string `json:"jwt"`
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
func CreateSession(ctx *gin.Context) {
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

	jwt := ""

	resBody := CreateSessionResBody{
		JWT: jwt,
	}

	ctx.JSON(http.StatusOK, resBody)

}
