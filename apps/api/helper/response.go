package helper

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{Message: message}
}

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{Message: message, Code: code}
}
