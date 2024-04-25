package domain

import "errors"

var (
	ErrAmountExceed = errors.New("amount exceed the limit")
)
