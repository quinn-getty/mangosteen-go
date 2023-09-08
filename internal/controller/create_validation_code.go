package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var body struct {
		Email string
	}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, "参数错误")
	}
	log.Println(body)
}
