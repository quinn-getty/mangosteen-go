package controller

import (
	"encoding/json"
	"log"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	meController := MeController{}
	meController.RegisterRouter(apiV1)

	// 不传header
	req, _ := http.NewRequest("GET", "/api/v1/me", strings.NewReader(""))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// 错误header
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/me", strings.NewReader(""))
	req.Header = http.Header{
		Authorization: []string{"xxxxxxxxx"},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// 空 header Authorization
	w = httptest.NewRecorder()
	jwtString := ""
	req, _ = http.NewRequest("GET", "/api/v1/me", strings.NewReader(""))
	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// user_id 错误
	w = httptest.NewRecorder()
	jwtString, err := jwt_helper.GenerateJWT(9999999)
	if err != nil {
		log.Println(err)
	}
	req, _ = http.NewRequest("GET", "/api/v1/me", strings.NewReader(""))
	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// 加密其他信息
	w = httptest.NewRecorder()
	jwtString, err = jwt_helper.GenerateAnyObj(jwt_helper.MapClaims{})
	if err != nil {
		log.Println(err)
	}
	req, _ = http.NewRequest("GET", "/api/v1/me", strings.NewReader(""))
	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// header Authorization 正确
	w = httptest.NewRecorder()
	email := "xxxxx@xxxxx.com"
	// 提前插入到数据库
	user, err := q.FindUserByEmail(database.DBCtx, email)
	if err != nil {
		user, err = q.CreateUser(database.DBCtx, email)
		if err != nil {
			log.Println("创建失败")
		}
	}
	jwtString, err = jwt_helper.GenerateJWT(int(user.ID))
	if err != nil {
		log.Println(err)
	}
	req, _ = http.NewRequest("GET", "/api/v1/me", strings.NewReader(""))
	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	log.Println(w)
	assert.Equal(t, http.StatusOK, w.Code)

	bodyStr := w.Body.String()
	log.Println(bodyStr)
	resUser := GetMeResBody{}
	json.Unmarshal([]byte(bodyStr), &resUser)
	log.Println(resUser.Resourse.ID)
	assert.Equal(t, resUser.Resourse.ID, user.ID)
	assert.Equal(t, resUser.Resourse.Email, user.Email)

}
