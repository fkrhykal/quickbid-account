package validation

import (
	"context"
	"fmt"
)

type Validator[T any] func(ctx context.Context, v T) error

type ValidationError struct {
	Detail Detail
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %+v", v.Detail)
}

type Detail map[string]string

func (d Detail) Add(field string, err error) {
	d[field] = err.Error()
}
