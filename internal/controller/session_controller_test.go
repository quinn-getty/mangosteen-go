package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)

	apiV1 := r.Group("/api/v1")
	sessionController := SessionController{}
	sessionController.RegisterRouter(apiV1)

	email := "xxxxx@xxxxx.com"
	code := "888888"

	// 提前插入到数据库
	user, err := q.FindUserByEmail(database.DBCtx, email)
	if err != nil {
		user, err = q.CreateUser(database.DBCtx, email)
		if err != nil {
			log.Println("创建失败")
		}
	}

	_, err = q.CreateValidationCode(database.DBCtx, queries.CreateValidationCodeParams{
		Email: email,
		Code:  code,
	})
	if err != nil {
		log.Fatalln(err)
	}

	j := gin.H{
		"email": email,
		"code":  code,
	}

	bytes, _ := json.Marshal(j)

	req, _ := http.NewRequest("POST", "/api/v1/session", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	responsBody := CreateSessionResBody{}
	if err = json.Unmarshal(w.Body.Bytes(), &responsBody); err != nil {
		log.Fatalln(err)
		t.Error("没有返回jwt")
	}

	fmt.Println(responsBody.JWT)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, user.ID, responsBody.UserId)
}
