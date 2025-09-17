package conf

import (
	"bytes"
	"net/http"
	s "strings"

	"github.com/gin-gonic/gin"
)

// StandardResponse defines the structure for all API responses
type StandardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Custom response writer to capture the response
type CustomResponseWriter struct {
	gin.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

func (r *CustomResponseWriter) Write(b []byte) (int, error) {
	return r.Body.Write(b)
}

func (r *CustomResponseWriter) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
}

func (r *CustomResponseWriter) Status() int {
	if r.StatusCode == 0 {
		return 200 // Default status
	}
	return r.StatusCode
}

// Enhanced context with custom response methods
type EnhancedContext struct {
	*gin.Context
}

// Custom response methods that bypass middleware wrapping
func (ec *EnhancedContext) StandardJSON(statusCode int, message string, data any) {
	response := StandardResponse{
		Message: message,
		Success: statusCode >= 200 && statusCode < 300,
		Data:    data,
	}
	ec.Context.JSON(statusCode, response)
}

func (ec *EnhancedContext) Success(data any) {
	ec.StandardJSON(http.StatusOK, "Request Successful", data)
}

func (ec *EnhancedContext) SuccessWithMessage(message string, data any) {
	ec.StandardJSON(http.StatusOK, message, data)
}

func (ec *EnhancedContext) SuccessWithMessageAndNoData(message string) {
	ec.StandardJSON(http.StatusOK, message, nil)
}

func (ec *EnhancedContext) SuccessWithNoMessageAndNoData() {
	ec.StandardJSON(http.StatusOK, "Request Success", nil)
}

func (ec *EnhancedContext) BadRequest(data any) {
	ec.StandardJSON(http.StatusBadRequest, "Request Failed", data)
}

func (ec *EnhancedContext) BadRequestWithMessage(message string, data any) {
	ec.StandardJSON(http.StatusBadRequest, message, data)
}

func (ec *EnhancedContext) BadRequestWithMessageAndNoData(message string) {
	ec.StandardJSON(http.StatusBadRequest, message, nil)
}

func (ec *EnhancedContext) BadRequestWithNoMessageAndNoData() {
	ec.StandardJSON(http.StatusBadRequest, "Request Failed", nil)
}

type validationError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

func parseValidationErrors(input string) []validationError {
	lines := s.Split(input, "\n")
	var result []validationError

	for _, line := range lines {
		// Split only at the first occurrence of " Error:"
		parts := s.SplitN(line, " Error:", 2)
		if len(parts) == 2 {
			result = append(result, validationError{
				Key:   s.Replace(s.Split(s.TrimSpace(parts[0]), " ")[1], "'", "", 2),
				Error: s.TrimSpace(parts[1]),
			})
		}
	}

	return result
}

func (ec *EnhancedContext) ValidationError(errorDetails string) {
	errors := parseValidationErrors(errorDetails)
	if len(errors) != 0 {
		ec.StandardJSON(http.StatusBadRequest, "Validation Failed", map[string]any{"details": errors})
	} else {
		ec.StandardJSON(http.StatusBadRequest, "Validation Failed", map[string]string{"details": errorDetails})
	}
}
