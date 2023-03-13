package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
)

func (h *Handler) logIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "GET":
		if r.Context().Value(CtxReqUserIdKey) != "Guest" {
			renderResponse(w, "allowed", http.StatusAccepted)
			return
		}
	case "POST":
		var user *entity.User
		data, err := io.ReadAll(r.Body)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: logIn: reading body")
			renderErrorResponse(w, "internal error", err)
			return
		}
		err = json.Unmarshal(data, &user)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: logIn: unmarshal error")
			renderErrorResponse(w, "internal error", err)
			return
		}
		cookie, err := h.UseCases.Users.SetCookie(r.Context(), user)
		if err != nil {
			renderErrorResponse(w, "invalid mail or password", err)
			return
		}
		http.SetCookie(w, cookie)
	}
}
