package controller

import (
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestCreateItem(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	itemController := ItemController{}
	itemController.RegisterRouter(apiV1)

	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/item",
		strings.NewReader(`{
			"amount": 100,
			"kind": "expenses",
			"happenedAt": "2020-01-01T00:00:00Z",
			"tagIds": [1, 2, 3]
		}`),
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	bodyStr := w.Body.String()
	resItem := CreateItemRes{}
	json.Unmarshal([]byte(bodyStr), &resItem)
	assert.Equal(t, resItem.Resource.UserID, user.ID)
	assert.Equal(t, int32(100), resItem.Resource.Amount)

}

func TestCreateItemWithError(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	itemController := ItemController{}
	itemController.RegisterRouter(apiV1)

	_, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/item",
		strings.NewReader(`{
			"kind": "expenses",
			"happenedAt": "2020-01-01T00:00:00Z",
			"tagIds": [1, 2, 3]
		}`),
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

}

func TestListItem(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	itemController := ItemController{}
	itemController.RegisterRouter(apiV1)

	_, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"get",
		"/api/v1/item",
		nil,
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
