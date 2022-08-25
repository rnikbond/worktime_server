package service

import (
	"errors"
	"net/http"
)

type ErrService struct {
	msg string
}

func NewErr(msg string) ErrService {
	return ErrService{
		msg: msg,
	}
}

func (err ErrService) Error() string {
	return err.msg
}

var ( // Login errors
	ErrUserUnauthorized = NewErr("user already exists")
)

func ToHTTP(err error) int {

	var serviceErr ErrService
	if !errors.As(err, &serviceErr) {
		return http.StatusInternalServerError
	}

	switch serviceErr {
	case ErrUserUnauthorized:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}
