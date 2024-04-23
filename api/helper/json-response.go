package helper

import (
	"encoding/json"
	"net/http"
	"status/internal/apiError"
)

func ReturnJsonResponse(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		apiError.ReturnApiError(writer, apiError.GetApiError(500, "could not marshal json"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
	return
}
