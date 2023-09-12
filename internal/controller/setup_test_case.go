package controller

import (
	"mangosteen/config/queries"
	"mangosteen/internal"
	"mangosteen/internal/database"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
* @desc init q w r
* @return func(t *testing.T)
 */
func setupTestCase(t *testing.T) (*queries.Queries, *httptest.ResponseRecorder, *gin.Engine, func(t *testing.T)) {
	r := gin.Default()
	internal.InitRouter(r)

	w := httptest.NewRecorder()
	q := database.NewQuery()

	return q, w, r, func(t *testing.T) {
		database.Close()
	}
}
