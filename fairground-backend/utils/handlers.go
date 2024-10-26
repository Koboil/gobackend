package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sends a JSON response with a status code and message.
func JSONResponse(w http.ResponseWriter, statusCode int, message interface{}, customMessage ...string) {

	// Step: 1 - Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Step: 2 - Encode the response message
	if errMsg, ok := message.(error); ok {

		// Step: 2a - If the message is an error, encode the error message
		errorResponse := map[string]string{
			"error": errMsg.Error(),
		}
		if len(customMessage) > 0 {
			errorResponse["error"] = customMessage[0]
		}
		json.NewEncoder(w).Encode(errorResponse)
	} else {

		// Step: 2b - Check if the message contains the "token" field, regardless of the map type
		if messageMapStr, ok := message.(map[string]string); ok {
			if token, exists := messageMapStr["token"]; exists && token != "" {
				// If message is a map[string]string and contains a valid "token", encode it as-is
				json.NewEncoder(w).Encode(messageMapStr)
				return
			}
		}

		// Step: 2c - Otherwise, wrap the response in a success message
		successResponse := map[string]interface{}{
			"success": message,
		}

		json.NewEncoder(w).Encode(successResponse)

	}
}

// JSONDataResponse sends a JSON response with a status code and data.
func JSONDataResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// add "data" key to the response
	data = map[string]interface{}{
		"data": data,
	}

	json.NewEncoder(w).Encode(data)

}
