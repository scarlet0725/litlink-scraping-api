package adapter

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

var InternalServerError = ErrorResponse{
	Ok:      false,
	Message: "Internal Server Error",
	Error:   "internal_server_error",
}

var BadRequest = ErrorResponse{
	Ok:      false,
	Message: "Bad Request",
	Error:   "bad_request",
}

func NewErrorResponse(message string, err error) ErrorResponse {
	return ErrorResponse{
		Ok:      false,
		Message: message,
		Error:   err.Error(),
	}
}
