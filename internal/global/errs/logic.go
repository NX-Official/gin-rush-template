package errs

import (
	"errors"
	"fmt"
)

const ErrorContextKey = "error"

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
	Origin  string `json:"origin"`
}

func newError(code int32, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
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

// WithOrigin 向前端返回用于调试的原始错误（仅限 config.DebugMode）
func (e *Error) WithOrigin(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Origin:  fmt.Sprintf("%+v", err),
	}
}

// WithTips 向前端返回额外的提示信息（config.ReleaseMode 也可见）
func (e *Error) WithTips(details ...string) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message + " " + fmt.Sprintf("%v", details),
	}
}
