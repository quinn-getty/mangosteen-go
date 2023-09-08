package controller_test

import (
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/v1/create_validation_code", strings.NewReader(`{"email": "quinn.getty@qq.com"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pang", w.Body.String())
}
