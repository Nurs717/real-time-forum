package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "POST":
		var user *entity.User
		data, err := io.ReadAll(r.Body)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: signUp: read body req")
			renderErrorResponse(w, "internal error", err)
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: signUp: unmarshal body")
			renderErrorResponse(w, "internal error", err)
			return
		}

		if err = h.UseCases.Users.NewUser(r.Context(), user); err != nil {
			renderErrorResponse(w, "creat user", err)
			return
		}
		log.Printf("user: %s created\n", user.UserName)
		renderResponse(w, nil, http.StatusCreated)
	}
}
