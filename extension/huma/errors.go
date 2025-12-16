package huma

import "github.com/danielgtaylor/huma/v2"

func Error409Conflict(msg string) huma.StatusError {
	return huma.Error409Conflict(msg)
}

func Error500InternalServerError() huma.StatusError {
	return huma.Error500InternalServerError("Internal server error")
}
