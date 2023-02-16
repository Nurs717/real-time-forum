package errors

import "errors"

// ErrEmptyRegisterData invalid entity
var ErrEmptyRegisterData = errors.New("missing required field when creating user")

// ErrEmptyPost empty post
var ErrEmptyPost = errors.New("the post is empty")

// ErrMailNotExist when loging in there is no user with that email
var ErrMailNotExist = errors.New("there is no user with this mail")

// ErrWrongPassword password of User is wrong while log in
var ErrWrongPassword = errors.New("password of User is wrong while log in")

// ErrInvalidPassword missing required letter when registering
var ErrInvalidPassword = errors.New("invalid password password when trying to register, missing Upper or Number or len less than 7")

// ErrTokenInvalid token is not found when trying to get authorization
var ErrTokenInvalid = errors.New("token not found in db")

// ErrEmailInvalid when email string invalid format
var ErrEmailInvalid = errors.New("invalid email format")
