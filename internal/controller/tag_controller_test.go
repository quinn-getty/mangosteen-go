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

	"github.com/stretchr/testify/assert"
)

func TestCreateTagWithSuccess(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	tagController := TagController{}
	tagController.RegisterRouter(apiV1)

	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/tag",
		strings.NewReader(`{
			"name": "通勤",
			"sign": "🚗"
		}`),
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	bodyStr := w.Body.String()
	resTag := CreateTagRes{}
	json.Unmarshal([]byte(bodyStr), &resTag)
	assert.Equal(t, resTag.Resource.UserID, user.ID)
	assert.Equal(t, resTag.Resource.Name, "通勤")
	assert.Equal(t, resTag.Resource.Sign, "🚗")
}

func TestCreateTagWithError(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	tagController := TagController{}
	tagController.RegisterRouter(apiV1)

	_, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/tag",
		strings.NewReader(`{
			"name": "",
			"sign": "🚗"
		}`),
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestListTagWithError(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	tagController := TagController{}
	tagController.RegisterRouter(apiV1)

	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	if err := q.DeleteUserAllTag(database.DBCtx, user.ID); err != nil {
		log.Fatalln(err)
	}

	q.CreateTag(database.DBCtx, queries.CreateTagParams{
		UserID: user.ID,
		Name:   "string",
		Sign:   "string",
	})

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/tag", nil,
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	bodyStr := w.Body.String()
	resTag := TagListRes{}
	json.Unmarshal([]byte(bodyStr), &resTag)
	assert.Equal(t, len(resTag.Resource), int(1))
}

func TestUpdateTagWithSuccess(t *testing.T) {
	q, w, r, teardownTest := setupTestCase(t)
	defer teardownTest(t)
	apiV1 := r.Group("/api/v1")
	tagController := TagController{}
	tagController.RegisterRouter(apiV1)

	user, jwtString, err := getUsereAndJwt(q)
	if err != nil {
		log.Println(err)
	}

	tag, err := q.CreateTag(database.DBCtx, queries.CreateTagParams{
		UserID: user.ID,
		Name:   "通勤",
		Sign:   "🚗",
	})
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest(
		"PATCH",
		"/api/v1/tag",
		strings.NewReader(fmt.Sprintf(`{
			"id": %d,
			"name": "吃饭",
			"sign": "🍚"
		}`, tag.ID)),
	)

	req.Header = http.Header{
		Authorization: []string{"Bearer " + jwtString},
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	bodyStr := w.Body.String()
	resTag := TagUpdateRes{}
	json.Unmarshal([]byte(bodyStr), &resTag)
	assert.Equal(t, resTag.Resource.ID, tag.ID)
	assert.Equal(t, resTag.Resource.Name, "吃饭")
	assert.Equal(t, resTag.Resource.Sign, "🍚")
}

func TestDeleteTagWithSuccess(t *testing.T) {
	// q, w, r, teardownTest := setupTestCase(t)
	// defer teardownTest(t)
	// apiV1 := r.Group("/api/v1")
	// tagController := TagController{}
	// tagController.RegisterRouter(apiV1)

	// user, jwtString, err := getUsereAndJwt(q)
	// if err != nil {
	// 	log.Println(err)
	// }

	// tag, err := q.CreateTag(database.DBCtx, queries.CreateTagParams{
	// 	UserID: user.ID,
	// 	Name:   "通勤",
	// 	Sign:   "🚗",
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// req, _ := http.NewRequest(
	// 	"PATCH",
	// 	"/api/v1/tag",
	// 	strings.NewReader(fmt.Sprintf(`{
	// 		"id": %d,
	// 		"name": "要删除的",
	// 		"sign": "🍚"
	// 	}`, tag.ID)),
	// )

	// req.Header = http.Header{
	// 	Authorization: []string{"Bearer " + jwtString},
	// }
	// r.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)

	// bodyStr := w.Body.String()
	// resTag := TagUpdateRes{}
	// json.Unmarshal([]byte(bodyStr), &resTag)
	// assert.Equal(t, resTag.Resource.ID, tag.ID)
	// assert.Equal(t, resTag.Resource.Name, "吃饭")
	// assert.Equal(t, resTag.Resource.Sign, "🍚")
}
