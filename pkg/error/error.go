package merror

import (
	"fmt"

	"github.com/pkg/errors"
)

// 文章：https://learnku.com/go/t/33210
type ErrorType uint

const (
	NoType = ErrorType(iota)
	BadRequest
	NotFound
)

type errorPlus struct {
	errType ErrorType
	err     error
	context errorContext
}

type errorContext struct {
	Field, Message string
}

func (et errorPlus) Error() string {
	return et.err.Error()
}

// New 用于创建一个 errorPlus 对象
func (et ErrorType) New(msg string) error {
	return errorPlus{errType: et, err: errors.New(msg)}
}

// Newf 用于格式化创建一个 errorPlus 对象
func (et ErrorType) Newf(format string, args ...interface{}) error {
	newErr := fmt.Errorf(format, args...)
	return errorPlus{errType: et, err: newErr}
}

// Wrap 方法新建一个封装错误
func (et ErrorType) Wrap(err error, msg string) error {
	return et.Wrapf(err, msg)
}

// Wrapf 格式化消息创建新的封装错误，继承原来错误的上下文
func (et ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return errorPlus{errType: et, err: errors.Wrapf(err, msg, args)}
}

// New 用于创建一个 errorPlus 对象
func New(msg string) error {
	return errorPlus{errType: NoType, err: errors.New(msg)}
}

// Newf 用于格式化创建一个 errorPlus 对象
func Newf(format string, args ...interface{}) error {
	newErr := fmt.Errorf(format, args...)
	return errorPlus{errType: NoType, err: newErr}
}

// Wrap 方法新建一个封装错误
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Wrapf 格式化消息创建新的封装错误，继承原来错误的上下文
func Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args)

	if plus, ok := err.(errorPlus); ok {
		return errorPlus{
			errType: plus.errType,
			err:     newErr,
			context: plus.context,
		}
	}
	return errorPlus{errType: NoType, err: newErr}
}

// AddErrorContext 方法为错误添加上下文
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}

	if plus, ok := err.(errorPlus); ok {
		return errorPlus{
			errType: plus.errType,
			err:     plus.err,
			context: context,
		}
	}
	return errorPlus{
		errType: NoType,
		err:     err,
		context: context,
	}
}

func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if plus, ok := err.(errorPlus); ok || plus.context != emptyContext {
		return map[string]string{
			"field":   plus.context.Field,
			"message": plus.context.Message,
		}
	}
	return nil
}

func GetType(err error) ErrorType {
	if plus, ok := err.(errorPlus); ok {
		return plus.errType
	}
	return NoType
}
