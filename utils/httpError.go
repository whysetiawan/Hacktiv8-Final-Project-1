package utils

type HttpError struct {
	Message interface{} `json:"message"`
	Trace   interface{} `json:"stack_trace"`
}

func NewHttpError(message interface{}, trace interface{}) HttpError {
	return HttpError{
		Message: message,
		Trace:   trace,
	}
}
