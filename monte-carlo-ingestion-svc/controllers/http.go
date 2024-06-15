package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime/debug"

	"monte-carlo-ingestion/domain"
	customError "monte-carlo-ingestion/errors"
	"monte-carlo-ingestion/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ParsedRequest holds the request metadata extracted from gin context
type ParsedRequest struct {
	Method    string `json:"method"`
	Route     string `json:"route"`
	IP        string `json:"ip,omitempty"`
	Protocol  string `json:"protocol"`
	Referer   string `json:"referer,omitempty"`
	UserAgent string `json:"userAgent"`
}

// ParsedResponse holds the data extracted from the response
type ParsedResponse struct {
	Size       int     `json:"size"`
	StatusCode int     `json:"statusCode"`
	TimeTaken  float64 `json:"timeTaken"`
}

// ParsedError holds the data extracted from error
type ParsedError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func parseRequest(c *gin.Context) ParsedRequest {
	return ParsedRequest{
		Method:    c.Request.Method,
		Route:     c.Request.RequestURI,
		IP:        c.ClientIP(),
		Protocol:  c.Request.Proto,
		Referer:   c.Request.Header.Get("Referer"),
		UserAgent: c.Request.Header.Get("User-Agent"),
	}
}

func parseResponse(c *gin.Context) ParsedResponse {
	return ParsedResponse{
		Size:       c.Writer.Size(),
		StatusCode: c.Writer.Status(),
	}
}

func parseError(err customError.BaseError) ParsedError {
	return ParsedError{
		Message: err.GetMsg(),
		Details: err.GetErrorContext(),
	}
}

type HTTPRequestData struct {
	Req ParsedRequest `json:"req"`
}

type HTTPResponseData struct {
	Req ParsedRequest  `json:"req"`
	Res ParsedResponse `json:"res"`
}

type HTTPResponseWithErrorData struct {
	Req ParsedRequest  `json:"req"`
	Res ParsedResponse `json:"res"`
	Err ParsedError    `json:"err"`
}

// LogRequest helps log the request data tied to a gin HTTP request
func LogRequest(c *gin.Context) {
	logger.Info("http", zap.String("type", "http-request"), zap.Any("data", HTTPRequestData{
		Req: parseRequest(c),
	}))
}

// LogResponse helps log the response tied to a gin HTTP request
func LogResponse(c *gin.Context) {
	logger.Info("http", zap.String("type", "http-response"), zap.Any("data", HTTPResponseData{
		Req: parseRequest(c),
		Res: parseResponse(c),
	}))
}

// LogError helps log the given error and associated metadata for the gin HTTP request
func LogError(c *gin.Context, err customError.BaseError) {
	logger.Error("http", zap.String("type", "http-response"), zap.Any("data", HTTPResponseWithErrorData{
		Req: parseRequest(c),
		Res: parseResponse(c),
		Err: parseError(err),
	}))
}

// HTTPController is a wrapper over controllers to convert them into a HTTP controller
func HTTPController(controller Controller) func(c *gin.Context) {
	return func(c *gin.Context) {
		var controllerRequest *domain.Request

		defer func() {
			if err := recover(); err != nil {
				HandleHTTPError(c, customError.NewInternalServerError(
					customError.WithError(err),
					customError.WithStackTrace(string(debug.Stack())),
					customError.WithErrorContext(controllerRequest),
				))
			}
		}()

		LogRequest(c)

		controllerRequest = controller.NewRequest()

		if controllerRequest.Body != nil {
			// Read the raw body
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				// Failed to read the body
				logger.Error("Error reading request body", zap.Error(err))
				// Handle error for failing to read body
				HandleHTTPError(c, customError.NewBadRequestError(
					customError.WithMsg("Failed to read request body"),
					customError.WithUserContext(map[string]interface{}{
						"error": err.Error(),
					}),
				))
				return
			}

			if c.Request.Header.Get("Content-Type") == "application/json" {
				if err := json.Unmarshal(bodyBytes, controllerRequest.Body); err != nil {
					// Handle the JSON unmarshal error
					HandleHTTPError(c, customError.NewBadRequestError(
						customError.WithError(err),
						customError.WithUserContext(map[string]interface{}{
							"body": string(bodyBytes),
						}),
					))
					return
				}
			}
		}

		if controllerRequest.QueryParams != nil {
			if err := c.ShouldBindQuery(controllerRequest.QueryParams); err != nil {
				// Handle the query binding error
				HandleHTTPError(c, customError.NewBadRequestError(
					customError.WithError(err),
					customError.WithUserContext(map[string]interface{}{
						"query": c.Request.URL.Query(),
					}),
				))
				return
			}
		}

		if controllerRequest.PathParams != nil {
			if err := c.ShouldBindUri(controllerRequest.PathParams); err != nil {
				// Handle the URI binding error
				HandleHTTPError(c, customError.NewBadRequestError(
					customError.WithError(err),
					customError.WithUserContext(map[string]interface{}{
						"params": c.Params,
					}),
				))
				return
			}
		}

		if controllerRequest.Headers != nil {
			if err := c.ShouldBindHeader(controllerRequest.Headers); err != nil {
				// Handle the header binding error
				HandleHTTPError(c, customError.NewBadRequestError(
					customError.WithError(err),
					customError.WithUserContext(map[string]interface{}{
						"headers": c.Request.Header,
					}),
				))
				return
			}
		}

		if err := controller.ValidateRequest(c, controllerRequest); err != nil {
			HandleHTTPError(c, customError.NewBadRequestError(
				customError.WithMsg(err.Error()),
				customError.WithErrorContext(controllerRequest),
			))
			return
		}

		controllerResponse, err := controller.Handler(c, controllerRequest)

		if err != nil {
			HandleHTTPError(c, err)
			return
		}

		c.JSON(http.StatusOK, domain.NewOkResponse(controllerResponse))
		LogResponse(c)
	}
}

// HandleHTTPError handles errors in http requests
func HandleHTTPError(c *gin.Context, err interface{}) {
	requestError := customError.HandleError(err)
	LogError(c, requestError)
	c.JSON(requestError.GetStatusCode(), domain.NewErrorResponse(map[string]interface{}{
		"context": requestError.GetUserContext(),
		"message": requestError.GetMsg(),
	}))
}

