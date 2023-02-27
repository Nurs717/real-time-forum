package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "POST":
		var user *entity.User
		data, err := io.ReadAll(r.Body)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: SignUp: read body req")
			renderErrorResponse(w, "internal error", err)
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: SignUp: unmarshal body")
			renderErrorResponse(w, "internal error", err)
			return
		}

		if err = h.UseCases.Users.NewUser(r.Context(), user); err != nil {
			renderErrorResponse(w, "creat user", err)
			return
		}
		fmt.Printf("user: %s created", user.UserName)
		renderResponse(w, nil, http.StatusCreated)
	}
}
