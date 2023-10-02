package api

import (
	ports "bnt/bnt-indent-service/internal/port"
)

type HTTPHandler struct {
	indentService ports.IndentService
}

func NewHTTPHandler(
	countryService ports.IndentService) *HTTPHandler {
	return &HTTPHandler{
		indentService: countryService,
	}
}
