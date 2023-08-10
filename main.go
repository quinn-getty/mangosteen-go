package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		log.Println(" test message")
		ctx.String(http.StatusOK, "pang")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
