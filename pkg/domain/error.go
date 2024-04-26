package domain

import (
	"errors"
	"fmt"
)

type Error struct {
	HttpCode int
	Code     string
	Message  string
	Inner    error
}

func (e Error) Error() string {
	return fmt.Sprint(e.Message)
}

var (
	ErrAmountExceed = errors.New("amount exceed the limit")
)
