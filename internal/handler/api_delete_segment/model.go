package api_delete_segment

type Request struct {
	SegmentName string `json:"segment_name" validate:"required"`
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
