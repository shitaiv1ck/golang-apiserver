package core_request

import (
	core_errors "apiserver/internal/core/errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidate(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		return err
	}

	requestValidate := validator.New()

	if err := requestValidate.Struct(dest); err != nil {
		return fmt.Errorf("%v: %w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}
