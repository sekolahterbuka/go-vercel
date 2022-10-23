package handler

import (
	"net/http"

	"github.com/sekolahkita/go-api/server/handler"
)

func UserDetail(w http.ResponseWriter, r *http.Request) {
	handler := handler.NewHandler()
	switch r.Method {

	case "GET":
		handler.Detail(w, r)

	default:
		http.Error(w, "method "+r.Method+" not allowed", http.StatusBadRequest)
	}

}
