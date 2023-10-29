package response

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewSuccess(data any) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func NewError(err error) Response {
	return Response{
		Success: false,
		Message: err.Error(),
	}
}
