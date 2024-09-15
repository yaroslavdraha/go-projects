package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func ParseAndValidate(r *http.Request, body interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return fmt.Errorf("invalid request payload: %v", err)
	}

	if err := validate.Struct(body); err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}

	return nil
}
