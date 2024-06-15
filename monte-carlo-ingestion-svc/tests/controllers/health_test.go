package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"monte-carlo-ingestion/controllers"
	"monte-carlo-ingestion/wire"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthController(t *testing.T) {
	healthController, err := wire.InitializaHealthControllerWithMocks()
	assert.Nil(t, err)

	controllerFunction := controllers.HTTPController(&healthController)

	req, err := http.NewRequest("GET", "/", strings.NewReader(``))
	assert.Nil(t, err)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = req

	blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	controllerFunction(c)

	responseBody := make(map[string]interface{})
	json.Unmarshal([]byte(blw.body.String()), &responseBody)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, true, responseBody["success"])
}
