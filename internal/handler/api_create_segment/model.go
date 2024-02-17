package api_create_segment

type Request struct {
	Slug string `json:"slug" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{Status: StatusOK}
}

func Error() Response {
	return Response{Status: StatusError}
}
