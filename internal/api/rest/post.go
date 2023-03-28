package rest

import (
	"fmt"
	"net/http"
)

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println(r.URL.Query().Get("id"))
		w.WriteHeader(204)
	}
}
