package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
* @desc init q w r
* @return func(t *testing.T)
 */
func setupTestCase(t *testing.T) (*queries.Queries, *httptest.ResponseRecorder, *gin.Engine, func(t *testing.T)) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	internal.InitRouter(r)

	w := httptest.NewRecorder()
	q := database.NewQuery()

	return q, w, r, func(t *testing.T) {
		database.Close()
	}
}

func getUsereAndJwt(q *queries.Queries) (queries.User, string, error) {
	email := "xxxxx@xxxxx.com"
	// user := queries.User{}
	// 提前插入到数据库
	user, err := q.FindUserByEmail(database.DBCtx, email)
	if err != nil {
		user, err = q.CreateUser(database.DBCtx, email)
		if err != nil {
			log.Println("创建失败")
			return user, "", err
		}
	}

	jwtString, err := jwt_helper.GenerateJWT(int(user.ID))
	return user, jwtString, err
}
