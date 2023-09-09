package controller_test

import (
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/controller"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	email := "quinnn.gao@gmail.com"
	code := "123456"
	r := router.New()
	w := httptest.NewRecorder()

	// 提前插入到数据库
	q := database.NewQuery()
	_, err := q.CreateValidationCode(database.DBCtx, queries.CreateValidationCodeParams{
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

	responsBody := controller.CreateSessionResBody{}
	if err = json.Unmarshal(w.Body.Bytes(), &responsBody); err != nil {
		log.Fatalln(err)
		t.Error("没有返回jwt")
	}
	log.Println(responsBody)

	assert.Equal(t, 200, w.Code)
}
