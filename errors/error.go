package errors

import "errors"

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrEmptyPost empty post
var ErrEmptyPost = errors.New("the post is empty")

//ErrMailNotExist when loging in there is no user with that email
var ErrMailNotExist = errors.New("there is no user with this mail")

//ErrWrongPasswor password of User is wrong while loging in
var ErrWrongPassword = errors.New("password of User is wrong while loging in")

//ErrTokenInvalid token is not found when trying to get autorizaton
var ErrTokenInvalid = errors.New("token not found in db")
