package middlewares

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
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

	contentType := c.Writer.Header().Get("Content-Type")
	isJSONResponse := strings.Contains(contentType, "application/json")

	// Get the captured response
	statusCode := c.Writer.Status()
	responseBody := w.Body.Bytes()

	// Only wrap if there's actual response content
	if len(responseBody) > 0 && isJSONResponse {
		// Check if response is already in standard format
		var existingResponse map[string]any
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
			var existingData any

			// Try to extract existing message and data if they exist
			if existingResponse != nil {
				if msg, ok := existingResponse["message"]; ok {
					if msgStr, ok := msg.(string); ok {
						existingMessage = msgStr
					}
				}

				if data, ok := existingResponse["data"]; ok {
					existingData = data
				} else {
					existingData = nil // No data field, set to nil
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
			var responseData any
			if err := json.Unmarshal(responseBody, &responseData); err != nil {
				// If it's not valid JSON, use as string
				wrappedResponse.Data = string(responseBody)
			} else {
				if existingMessage != "" {
					// Remove the message from the response data if it exists
					if dataMap, ok := responseData.(map[string]any); ok {
						delete(dataMap, "message")
						// Check if data is empty after removing message
						if len(dataMap) == 0 {
							responseData = nil
						}
					}
				}

				if existingData != nil {
					// If existing data is present, use it
					wrappedResponse.Data = existingData
				} else {
					// Otherwise, use the unmarshaled response data
					wrappedResponse.Data = responseData
				}
			}

			// Check if the response body has other fields that need to be preserved
			var finalResponseData = make(map[string]any)
			if existingData != nil {
				for key, value := range existingResponse {
					if key != "message" && key != "success" && key != "data" {
						finalResponseData[key] = value
					}
				}
			}

			// Marshal the wrapped response
			var err error
			if len(finalResponseData) == 0 {
				finalResponse, err = json.Marshal(wrappedResponse)
			} else {
				// Include other fields in the final response
				finalResponseData["message"] = wrappedResponse.Message
				finalResponseData["success"] = wrappedResponse.Success
				finalResponseData["data"] = wrappedResponse.Data
				finalResponse, err = json.Marshal(finalResponseData)
			}

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
