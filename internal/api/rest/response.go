package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"rtforum/internal/cerror"
)

type ErrorResponse struct {
	Msg  string `json:"error_message"`
	Type string `json:"error_type,omitempty"`
}

func renderErrorResponse(w http.ResponseWriter, msg string, err error) {
	resp := ErrorResponse{Msg: msg, Type: cerror.DefaultType}
	status := http.StatusInternalServerError

	var ierr *cerror.Error
	if !errors.As(err, &ierr) {
		resp.Msg = "internal error"
	} else {
		switch ierr.Code() {
		case cerror.ErrorCodeConflict:
			status = http.StatusConflict
			resp.Msg = ierr.Msg()
			resp.Type = ierr.Type()
			logMsg := fmt.Sprintf("repo: NewUser: exec db: %s", ierr.Msg())
			err = cerror.WrapErrorf(ierr.Unwrap(), cerror.ErrorCodeConflict, "", logMsg)
			errors.As(err, &ierr)
		case cerror.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest
			resp.Msg = "validation requirements missing"
			resp.Type = ierr.Type()
		case cerror.ErrorCodeUnauthorized:
			status = http.StatusUnauthorized
		case cerror.ErrorCodeInternal:
			resp.Msg = "internal error"
		}
	}

	log.Printf("%s\n", ierr.Error())

	renderResponse(w, resp, status)
}

func renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		log.Printf("rest: Marshal response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(content); err != nil {
		log.Printf("rest: Write to response: %v\n", err)
	}
}
