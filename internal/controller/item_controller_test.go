package controller

import (
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

// 测试分页
func TestListItemOfPage(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	// 生成路由
	itemController := ItemController{}
	itemController.RegisterRouter(r.Group("/api/v1"))

	// 获取测试账户jwt
	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/item",
		nil,
	)

	// 清空测试用户的数据
	err = q.DeleteItem(database.DBCtx, user.ID)
	if err != nil {
		log.Fatal("清库", err)
	}

	// 构造用户的数据
	for i := 0; i < 5; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindInCome,
			HappenedAt: time.Now(),
			TagIds:     []int32{1, 2, 3},
		})
	}
	for i := 0; i < 5; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now().AddDate(0, 0, -1),
			TagIds:     []int32{1, 2, 3},
		})
	}
	for i := 0; i < 5; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindInCome,
			HappenedAt: time.Now().AddDate(0, 0, -2),
			TagIds:     []int32{1, 2, 3},
		})
	}

	// 测试 分页是否正常
	query := url.Values{}
	query.Add("size", "10")
	query.Add("current", "2")

	req.URL.RawQuery = query.Encode()

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	res := ItemGetListRes{}
	bodyStr := w.Body.String()
	json.Unmarshal([]byte(bodyStr), &res)

	assert.Equal(t, len(res.Resourses), 5)
	assert.Equal(t, res.Expenses, int32(5000))
	assert.Equal(t, res.Income, int32(10000))
}

func TestListItemOfQuery(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	// 生成路由
	itemController := ItemController{}
	itemController.RegisterRouter(r.Group("/api/v1"))

	// 获取测试账户jwt
	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/item",
		nil,
	)

	// 清空测试用户的数据
	err = q.DeleteItem(database.DBCtx, user.ID)
	if err != nil {
		log.Fatal("清库", err)
	}

	// 构造用户的数据
	for i := 0; i < 3; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now(),
			TagIds:     []int32{1},
		})
	}

	for i := 0; i < 3; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now().AddDate(0, 0, -2),
			TagIds:     []int32{2},
		})
	}

	for i := 0; i < 3; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now().AddDate(0, 0, -8),
			TagIds:     []int32{3},
		})
	}

	// 测试 分页是否正常
	query := url.Values{}
	query.Add("size", "10")
	query.Add("current", "1")
	query.Add("happenedAtBegin", time.Now().AddDate(0, 0, -2).Format(time.DateTime))
	query.Add("happenedAtEnd", time.Now().Format(time.DateTime))

	req.URL.RawQuery = query.Encode()
	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	res := ItemGetListRes{}
	bodyStr := w.Body.String()
	log.Println(res)
	json.Unmarshal([]byte(bodyStr), &res)

	assert.Equal(t, len(res.Resourses), int(3))
	assert.Equal(t, 1000*3, int(res.Expenses))
	assert.Equal(t, 0, int(res.Income))
}
func TestItemGetSummary(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	// 生成路由
	itemController := ItemController{}
	itemController.RegisterRouter(r.Group("/api/v1"))

	// 获取测试账户jwt
	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/item/summary",
		nil,
	)

	// 清空测试用户的数据
	err = q.DeleteItem(database.DBCtx, user.ID)
	if err != nil {
		log.Fatal("清库", err)
	}

	// 构造用户的数据
	for i := 0; i < 3; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now(),
			TagIds:     []int32{1},
		})
	}

	for i := 0; i < 3; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now().AddDate(0, 0, -2),
			TagIds:     []int32{2},
		})
	}

	for i := 0; i < 3; i++ {
		q.CreateItem(database.DBCtx, queries.CreateItemParams{
			UserID:     user.ID,
			Amount:     1000,
			Kind:       queries.KindExpenses,
			HappenedAt: time.Now().AddDate(0, 0, -8),
			TagIds:     []int32{3},
		})
	}

	// 测试 分页是否正常
	query := url.Values{}
	// query.Add("size", "10")
	// query.Add("current", "1")
	query.Add("happenedAtBegin", time.Now().AddDate(0, 0, -2).Format(time.DateTime))
	query.Add("happenedAtEnd", time.Now().Format(time.DateTime))

	req.URL.RawQuery = query.Encode()
	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// res := ItemGetListRes{}
	// bodyStr := w.Body.String()
	// log.Println(res)
	// json.Unmarshal([]byte(bodyStr), &res)

	// assert.Equal(t, len(res.Resourses), int(3))
	// assert.Equal(t, 1000*3, int(res.Expenses))
	// assert.Equal(t, 0, int(res.Income))
}
