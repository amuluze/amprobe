// Package errors
// Date: 2024/3/6 13:15
// Author: Amu
// Description:
package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	Is           = errors.Is
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	UnWrap       = errors.Unwrap
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	ErrInternalServer = New500Errors("internal service error")
	ErrBadParameter   = New400Errors("parameter parse error")
	ErrBadRequest     = New400Errors("bad request")
	ErrInvalidParent  = New400Errors("not found parent node")
	ErrUserDisable    = New400Errors("user forbidden")
)

type Errors struct {
	Message string // 错误消息
	Status  int    // 响应状态码
	Err     error  // 响应错误
}

func (r *Errors) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return r.Message
}

func NewErrors(status int, msg string, args ...interface{}) error {
	res := &Errors{
		Status:  status,
		Message: fmt.Sprintf(msg, args...),
	}
	return res
}

func New400Errors(msg string, args ...interface{}) error {
	return NewErrors(400, msg, args...)
}

func New500Errors(msg string, args ...interface{}) error {
	return NewErrors(500, msg, args...)
}

func UnWrapError(err error) *Errors {
	if v, ok := err.(*Errors); ok {
		return v
	}
	return nil
}

func WrapError(err error, status int, msg string, args ...interface{}) error {
	res := &Errors{
		Message: fmt.Sprintf(msg, args...),
		Err:     err,
		Status:  status,
	}
	return res
}

func Wrap400Error(err error) error {
	return WrapError(err, 400, "invalid parameter")
}

func Wrap500Error(err error) error {
	return WrapError(err, 500, "internal service error")
}
