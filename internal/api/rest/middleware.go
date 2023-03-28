package rest

import (
	"context"
	"log"
	"net/http"
	"time"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Set-Cookie")
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

func (h *Handler) reqLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			log.Printf("%v, %v, %v", r.Method, r.URL.Path, time.Since(start))
		}(time.Now())
		next.ServeHTTP(w, r)
	}
}
