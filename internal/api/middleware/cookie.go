package middleware

import "net/http"

func CheckCookie(cookie *http.Cookie) bool {
	return false
}
