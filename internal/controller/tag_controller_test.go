package controller

import (
	"encoding/json"
	"log"
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
