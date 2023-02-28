package cerror

import (
	"errors"
	"fmt"
)

// ErrEmptyRegisterData invalid entity
const ErrEmptyRegisterData = "usecase: entity: Validate: missing required field"

// ErrEmptyPost empty post
var ErrEmptyPost = errors.New("the post is empty")

// ErrInvalidPassword missing required letter when registering
const ErrInvalidPassword = "usecase: entity: Validate: invalid password, missing Upper, Lower or Number letter, or length less than 8"

// ErrTokenInvalid token is not found when trying to get authorization
var ErrTokenInvalid = errors.New("token not found in db")

// ErrEmailInvalid when email string invalid format
const ErrEmailInvalid = "usecase: entity: Validate: invalid email format"

type Error struct {
	origin  error
	msg     string
	errType string
	code    ErrorCode
}

type ErrorCode uint

const (
	ErrorCodeInternal ErrorCode = iota
	ErrorCodeUnknown
	ErrorCodeNotFound
	ErrorCodeInvalidArgument
	ErrorCodeConflict
	ErrorCodeUnauthorized
)

const (
	UserType       = "user"
	MailType       = "mail"
	FormatPwdType  = "password_format"
	FormatMailType = "mail_format"
	DefaultType    = "default"
)

func WrapErrorf(orig error, code ErrorCode, errorType string, format string, a ...interface{}) error {
	return &Error{
		code:    code,
		origin:  orig,
		msg:     fmt.Sprintf(format, a...),
		errType: errorType,
	}
}

func NewErrorf(code ErrorCode, errType string, format string, a ...interface{}) error {
	return WrapErrorf(nil, code, errType, format, a...)
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Error() string {
	if e.origin != nil {
		return fmt.Sprintf("%s: %s", e.msg, e.origin.Error())
	}
	return e.msg
}

func (e *Error) Unwrap() error {
	return e.origin
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Type() string {
	return e.errType
}
