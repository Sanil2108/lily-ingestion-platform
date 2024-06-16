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

func TestBadRequest(t *testing.T) {
	ingestionController, err := wire.InitializaIngestionControllerWithMocks()
	assert.Nil(t, err)

	controllerFunction := controllers.HTTPController(&ingestionController)

	req, err := http.NewRequest("POST", "/", strings.NewReader(``))
	assert.Nil(t, err)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = req

	blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	controllerFunction(c)

	responseBody := make(map[string]interface{})
	json.Unmarshal([]byte(blw.body.String()), &responseBody)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestIngestionController(t *testing.T) {
	ingestionController, err := wire.InitializaIngestionControllerWithMocks()
	assert.Nil(t, err)

	controllerFunction := controllers.HTTPController(&ingestionController)

	req, err := http.NewRequest(
		"POST",
		"/_api/ingestion",
		strings.NewReader(`
			{
				"apiKey":"myApiKey",
				"userId":"myUserId",
				"tenantId":"myTenantId",
				"healthStatus": {
					"status":"ok",
					"tableName":"myTableName",
					"timestamp":1718473247
				}
			}`,
		),
	)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", "application/json")

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
