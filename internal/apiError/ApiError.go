package apiError

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetApiError(code int, message string) *ApiError {
	return &ApiError{Code: code, Message: message}
}

func ReturnApiError(writer http.ResponseWriter, apiError *ApiError) {
	writer.Header().Set("Content-Type", "application/json")
	jsonRes, err := json.Marshal(apiError)
	if err != nil {
		http.Error(writer, "Could not marshal json. Something is really off", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(apiError.Code)
	writer.Write(jsonRes)
	return
}
