package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
)

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: reading body")
			renderErrorResponse(w, "internal error", err)
		}
		err = json.Unmarshal(data, &user)
		fmt.Printf("user login: %+v\n", user)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: unmarshal error")
			renderErrorResponse(w, "internal error", err)
			return
		}
		cookie, err := h.UseCases.SetCookie(r.Context(), user)
		if err != nil {
			renderErrorResponse(w, "invalid mail or password", err)
			return
		}
		http.SetCookie(w, cookie)
	}
}
