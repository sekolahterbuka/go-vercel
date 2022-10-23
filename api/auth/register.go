package handler

import (
	"net/http"

	"github.com/sekolahkita/go-api/server/handler"
	"github.com/sekolahkita/go-api/server/service"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	service := service.NewAuthService(&service.AuthConfig{})
	handler := handler.NewAuthHandler(&handler.AuthConfig{
		AuthService: service,
	})

	switch r.Method {

	case "POST":
		handler.Register(w, r)
	case "GET":
		handler.Register(w, r)

	default:
		http.Error(w, "method "+r.Method+" not allowed", http.StatusBadRequest)
	}

}
