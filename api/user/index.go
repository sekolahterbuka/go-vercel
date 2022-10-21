package handler

import (
	"net/http"

	"github.com/sekolahkita/go-api/server/handler"
)

func ChiId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("post"))
	case "GET":
		handler.GetUser(w, r)
	case "PUT":
		w.Write([]byte("put"))
	case "DELETE":
		w.Write([]byte("delete"))
	default:
		http.Error(w, "method "+r.Method+" not allowed", http.StatusBadRequest)
	}

}
