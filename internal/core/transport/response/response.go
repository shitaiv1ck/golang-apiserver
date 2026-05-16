package core_response

import (
	core_errors "apiserver/internal/core/errors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ResponseHandler struct {
	wr http.ResponseWriter
}

func NewResponseHandler(wr http.ResponseWriter) *ResponseHandler {
	return &ResponseHandler{
		wr: wr,
	}
}

func (r *ResponseHandler) JsonResponse(responseBody any, statusCode int) {
	r.wr.WriteHeader(statusCode)

	if err := json.NewEncoder(r.wr).Encode(responseBody); err != nil {
		fmt.Println(err)
	}
}

func (r *ResponseHandler) ErrorResponse(msg string, err error) {
	errorResponse := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	statusCode := r.setErrorStatusCode(err)
	r.wr.WriteHeader(statusCode)

	if err := json.NewEncoder(r.wr).Encode(errorResponse); err != nil {
		fmt.Println(err)
	}
}

func (r *ResponseHandler) setErrorStatusCode(err error) int {
	if errors.Is(err, core_errors.ErrInvalidArgument) {
		return http.StatusBadRequest
	}

	if errors.Is(err, core_errors.ErrNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, core_errors.ErrInvalidPasswordOrEmail) {
		return http.StatusUnauthorized
	}

	if errors.Is(err, core_errors.ErrConflict) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}
