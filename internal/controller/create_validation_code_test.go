package controller

import (
	"mangosteen/internal/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	vaildationController := VaildationCodeController{}
	vaildationController.RegisterRouter(apiV1)

	viper.Set("email.smtp.host", "localhost")
	viper.Set("email.smtp.port", "1025")

	email := "quinnn.gao@gmail.com"
	q := database.NewQuery()
	w := httptest.NewRecorder()

	count1, _ := q.CountValidationCodes(database.DBCtx, email)
	req, _ := http.NewRequest("POST", "/api/v1/create_validation_code", strings.NewReader(`{"email": "quinnn.gao@gmail.com"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(database.DBCtx, email)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, count2-1, count1)
}
