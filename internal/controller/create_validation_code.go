package controller

import (
	"crypto/rand"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateValidationCodeBody struct {
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
	var body CreateValidationCodeBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Println("入参错误")
		ctx.String(http.StatusBadRequest, "参数错误")
		return
	}

	code, err := generateCode()
	if err != nil {
		log.Println("code生成失败")
		ctx.String(http.StatusInternalServerError, "发送失败")
	}

	q := database.NewQuery()
	item, err := q.CreateValidationCode(ctx, queries.CreateValidationCodeParams{
		Email: body.Email,
		Code:  code,
	})

	if err != nil {
		log.Println("保存code失败")
		ctx.String(http.StatusBadRequest, "发送失败")
		return
	}

	if err := email.SendValidationCode(item.Email, item.Code); err != nil {
		log.Println("发送邮件失败")
		ctx.String(http.StatusInternalServerError, "发送失败")
	}
	ctx.String(http.StatusOK, "发送成功")
}

func generateCode() (string, error) {
	len := 6
	b := make([]byte, len)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	digits := make([]byte, len)
	for i := range b {
		digits[i] = b[i]%10 + 48
	}

	return string(digits), nil

}
