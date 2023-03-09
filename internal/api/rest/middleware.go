package rest

import (
	"context"
	"log"
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

type CtxKey string

const (
	CtxReqIdKey CtxKey = "X-Request-Id"
	//CtxReqUsernameKey CtxUsernameKey = "X-Request-Username"
)

func (h *Handler) checkCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		CtxKey := "Guest"
		if err == nil {
			userID, err := h.UseCases.Users.IsCookieValid(r.Context(), cookie.Value)
			if err != nil {
				log.Printf("middleware: checkCookie: r.Cookie: %v", err.Error())
			}
			if userID != "" {
				CtxKey = userID
			}
		}

		ctx1 := context.WithValue(r.Context(), CtxReqIdKey, CtxKey)
		ctx2 := r.WithContext(ctx1)

		next.ServeHTTP(w, ctx2)
	}
}
