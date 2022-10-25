package util

import (
	"net/http"

	apperrors "github.com/sekolahkita/go-api/server/util/apperror"
)

func SetupResponse(w http.ResponseWriter, body []byte, statusCode int) {
	contentType := "aplication/json"
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		apperrors.NewInternal(w)
	}
}
