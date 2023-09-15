package controller

import (
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"
	"net/url"
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

func TestCreateItemWithSuccess(t *testing.T) {
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
	// 生成路由
	itemController := ItemController{}
	itemController.RegisterRouter(r.Group("/api/v1"))

	// 获取测试账户jwt
	_, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req1, _ := http.NewRequest(
		"GET",
		"/api/v1/item",
		nil,
	)

	// // 生成一个item
	// for i := 0; i < 5; i++ {
	// 	q.CreateItem(database.DBCtx, queries.CreateItemParams{
	// 		UserID:     user.ID,
	// 		Amount:     1000,
	// 		Kind:       "in_come",
	// 		HappenedAt: time.Now(),
	// 		TagIds:     []int32{1, 2, 3},
	// 	})
	// }

	// q.ListItem(database.DBCtx, queries.ListItemParams{
	// 	UserID: int32(user.ID),
	// 	// HappenedAt   time.Time `json:"happenedAt"`
	// 	// HappenedAt_2 time.Time `json:"happenedAt2"`
	// 	Offset: 0,
	// 	Limit:  5,
	// })

	query := url.Values{}
	query.Add("size", "5")
	query.Add("current", "1")

	req1.URL.RawQuery = query.Encode()

	req1.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req1)

	assert.Equal(t, http.StatusOK, w.Code)

	res := ItemGetListRes{}
	bodyStr := w.Body.String()
	json.Unmarshal([]byte(bodyStr), &res)
	log.Println(res.Resourses)
	assert.Equal(t, len(res.Resourses), 5)
}
