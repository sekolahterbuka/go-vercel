package handler

import (
	"net/http"

	"github.com/sekolahkita/go-api/server/utils"
)

func (h *handler) Me(w http.ResponseWriter, r *http.Request) {
	params := utils.GetQuery(r, "id")
	w.Write([]byte("Hello = " + params))
}

func (h *handler) Token(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Token"))
}
func (h *handler) Detail(w http.ResponseWriter, r *http.Request) {
	params := utils.GetQuery(r, "id")

	w.Write([]byte("Hello Detail " + params))
}
