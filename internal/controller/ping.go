package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	log.Println(" test message")
	ctx.String(http.StatusOK, "pang")
}
