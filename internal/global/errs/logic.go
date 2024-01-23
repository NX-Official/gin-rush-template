package errs

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
	Origin  string `json:"origin"`
	//LogInfo map[string]string `json:"-"`
}

func newError(code int32, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
		//LogInfo: make(map[string]string),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", e.Code, e.Message)
}

func (e *Error) Is(target error) bool {
	var t *Error
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func (e *Error) WithOrigin(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Origin:  fmt.Sprintf("%+v", err),
	}
}

func (e *Error) WithTips(details ...string) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message + " " + fmt.Sprintf("%v", details),
	}
}
