package core_utils

import (
	core_errors "apiserver/internal/core/errors"
	"fmt"
	"net/http"
)

func GetStringPathValue(r *http.Request, key string) (string, error) {
	value := r.PathValue(key)
	if value == "" {
		return "", fmt.Errorf("email can't be null: %w", core_errors.ErrInvalidArgument)
	}

	return value, nil
}
