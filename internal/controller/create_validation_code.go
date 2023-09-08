package controller

import (
	"log"
	"mangosteen/internal/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateValidationCodeReq struct {
	Email string `json:"email" binging:"required,email"`
}

// 发送验证码 godoc
// @Summary      发送验证码
// @Description  发送验证码
// @Tags         验证码
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /create_validation_code [post]
func CreateValidationCode(ctx *gin.Context) {
	log.Println("--------")
	var body = CreateValidationCodeReq{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, "参数错误")
	}

	var code = "123456"

	if err := email.SendValidationCode(body.Email, code); err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "发送失败")
	}

	log.Println(body)
}
