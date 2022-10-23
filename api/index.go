package handler

import (
	"net/http"

	"github.com/sekolahkita/go-api/server/handler"
)

func init() {
}

func Chi(w http.ResponseWriter, r *http.Request) {
	handler := handler.NewHandler()

	handler.Me(w, r)

}
