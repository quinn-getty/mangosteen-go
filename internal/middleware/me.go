package middleware

import (
	"fmt"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Me(whitePaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		for _, s := range whitePaths {
			if has := strings.HasPrefix(path, s); has {
				c.Next()
				return
			}
		}

		user, err := getMe(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("me", user)
		c.Next()
	}
}

func GetMe(c *gin.Context) (queries.User, bool) {
	me, _ := c.Get("me")
	user, ok := me.(queries.User)
	return user, ok
}

func getMe(ctx *gin.Context) (queries.User, error) {
	user := queries.User{}

	auth := ctx.GetHeader("Authorization")
	if len(auth) < 8 {
		log.Print("无效的JWT Authorization < 8", auth)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return user, fmt.Errorf("无效的JWT")
	}

	jwtString := auth[7:]
	token, err := jwt_helper.Parse(jwtString)
	if err != nil {
		log.Print("无效的JWT jwtString jwt_helper.Parse 失败", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return user, fmt.Errorf("无效的JWT")

	}

	m, ok := token.Claims.(jwt_helper.MapClaims)
	if !ok {
		log.Print("无效的JWT token.Claims 失败", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return user, fmt.Errorf("无效的JWT")

	}

	userIdF64, ok := m["user_id"].(float64)
	userId := int32(userIdF64)
	if !ok {
		log.Print("无效的JWT userIdStr 获取失败", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return user, fmt.Errorf("无效的JWT")

	}

	q := database.NewQuery()
	user, err = q.FindUserById(database.DBCtx, userId)
	if err != nil {
		log.Println(userId)
		log.Print("无效的JWT FindUserById ", err)
		ctx.String(http.StatusUnauthorized, "无效的JWT")
		return user, fmt.Errorf("无效的JWT")
	}
	return user, nil
}
