package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login rest:", r.Context().Value(CtxReqIdKey))
	if r.Context().Value(CtxReqIdKey) != "Guest" {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	// fmt.Println(r.Cookie("session"))
	switch r.Method {
	case "POST":
		var user *entity.User
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("error reading body %v\n", err)
		}
		err = json.Unmarshal(data, &user)
		fmt.Println("user login", user)
		if user == nil {
			return
		}
		if err != nil {
			log.Printf("error unmarshaling %v\n", err)
		}
		cookie, err := h.UseCases.SetCookie(r.Context(), user)
		if err != nil {
			log.Printf("Error occured in LogIn rest: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(cookie)
		http.SetCookie(w, cookie)
	}
}
