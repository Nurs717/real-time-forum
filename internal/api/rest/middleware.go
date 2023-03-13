package rest

import (
	"context"
	"net/http"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Set-Cookie")
		next.ServeHTTP(w, r)
	}
}

type CtxUserIdKey string
type CtxUsernameKey string

const (
	CtxReqUserIdKey   CtxUserIdKey   = "X-Request-Id"
	CtxReqUsernameKey CtxUsernameKey = "X-Request-Username"
)

func (h *Handler) checkCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		userIdKey := "Guest"
		usernameKey := ""
		if err == nil {
			userID, err := h.UseCases.Users.IsCookieValid(r.Context(), cookie.Value)
			if err != nil {
				renderErrorResponse(w, "middleware error", err)
			}
			if userID != "" {
				userIdKey = userID
				usernameKey, err = h.UseCases.Users.GetUserName(r.Context(), userID)
				if err != nil {
					renderErrorResponse(w, "middleware error", err)
				}
			}
		}

		ctx1 := context.WithValue(r.Context(), CtxReqUserIdKey, userIdKey)
		ctx1 = context.WithValue(ctx1, CtxReqUsernameKey, usernameKey)
		ctx2 := r.WithContext(ctx1)

		next.ServeHTTP(w, ctx2)
	}
}
