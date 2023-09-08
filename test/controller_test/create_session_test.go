package controller_test

import (
	"encoding/json"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()

	j := gin.H{
		"email": "quinnn.gao@gmail.com",
		"code":  "",
	}

	bytes, _ := json.Marshal(j)

	req, _ := http.NewRequest("POST", "/api/v1/session", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
