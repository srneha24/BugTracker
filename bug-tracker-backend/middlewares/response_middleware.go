package middlewares

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
)

func StandardResponseMiddleware(c *gin.Context) {
	// Create a custom response writer to capture the response
	w := &conf.CustomResponseWriter{
		ResponseWriter: c.Writer,
		Body:           bytes.NewBufferString(""),
		StatusCode:     0,
	}
	c.Writer = w

	// Process the request
	c.Next()

	// Get the captured response
	statusCode := c.Writer.Status()
	responseBody := w.Body.Bytes()

	// Only wrap if there's actual response content
	if len(responseBody) > 0 {
		// Check if response is already in standard format
		var existingResponse map[string]interface{}
		isStandardFormat := false

		if err := json.Unmarshal(responseBody, &existingResponse); err == nil {
			if _, hasMessage := existingResponse["message"]; hasMessage {
				if _, hasSuccess := existingResponse["success"]; hasSuccess {
					if _, hasData := existingResponse["data"]; hasData {
						// Already in standard format, don't wrap
						isStandardFormat = true
					}
				}
			}
		}

		var finalResponse []byte

		if isStandardFormat {
			// Use the original response as-is
			finalResponse = responseBody
		} else {
			// Wrap the response in standard format
			var wrappedResponse conf.StandardResponse
			var existingMessage string

			// Try to extract existing message if it exists
			if existingResponse != nil {
				if msg, ok := existingResponse["message"]; ok {
					if msgStr, ok := msg.(string); ok {
						existingMessage = msgStr
					}
				}
			}

			if statusCode >= 200 && statusCode < 300 {
				wrappedResponse.Success = true
				if existingMessage != "" {
					wrappedResponse.Message = existingMessage
				} else {
					wrappedResponse.Message = "Request Successful"
				}
			} else {
				wrappedResponse.Success = false
				if existingMessage != "" {
					wrappedResponse.Message = existingMessage
				} else {
					wrappedResponse.Message = "Request Failed"
				}
			}

			// Set the data
			if len(responseBody) > 0 {
				var responseData interface{}
				if err := json.Unmarshal(responseBody, &responseData); err != nil {
					// If it's not valid JSON, use as string
					wrappedResponse.Data = string(responseBody)
				} else {
					if existingMessage != "" {
						// Remove the message from the response data if it exists
						if dataMap, ok := responseData.(map[string]interface{}); ok {
							delete(dataMap, "message")
							// Check if data is empty after removing message
							if len(dataMap) == 0 {
								responseData = nil
							}
						}
					}
					wrappedResponse.Data = responseData
				}
			} else {
				wrappedResponse.Data = nil
			}

			// Marshal the wrapped response
			var err error
			finalResponse, err = json.Marshal(wrappedResponse)
			if err != nil {
				// Fallback to original response if marshaling fails
				finalResponse = responseBody
			}
		}

		// Set proper headers and write the final response
		w.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.ResponseWriter.WriteHeader(statusCode)
		w.ResponseWriter.Write(finalResponse)
	} else if len(responseBody) > 0 {
		// Non-JSON response, pass through as-is
		for key, values := range w.Header() {
			for _, value := range values {
				w.ResponseWriter.Header().Add(key, value)
			}
		}
		w.ResponseWriter.WriteHeader(statusCode)
		w.ResponseWriter.Write(responseBody)
	} else {
		// Empty response body
		w.ResponseWriter.WriteHeader(statusCode)
	}
}

// Middleware to inject enhanced context
func EnhancedContextMiddleware(c *gin.Context) {
	ec := &conf.EnhancedContext{Context: c}
	c.Set("enhanced", ec)
	c.Next()
}
