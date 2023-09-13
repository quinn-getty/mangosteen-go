package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
//	@Summary		测试连通性
//	@Description	测试连通性
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/ping [get]
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pang")
}
