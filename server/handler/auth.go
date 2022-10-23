package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/sekolahkita/go-api/server/model"
	"github.com/sekolahkita/go-api/server/utils"
)

type authHandler struct {
	AuthService model.AuthService
}
type AuthConfig struct {
	AuthService model.AuthService
}

func NewAuthHandler(c *AuthConfig) *authHandler {
	return &authHandler{
		AuthService: c.AuthService,
	}
}

// func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
// 	w.Header().Set("Content-Type", contentType)
// 	w.WriteHeader(statusCode)
// 	_, err := w.Write(body)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func (h *authHandler) reqSerializer(contentType string) model.RegisterSerializer {
// 	return &json.RegisterParams{}
// }

// func (h *authHandler) resSerializer(contentType string) model.AuthSerializer {
// 	return &json.Auth{}
// }

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := h.AuthService.Register(ctx, model.RegisterParams{Username: "sanja", Email: "email", Password: "pasword"})
	if err != nil {
		log.Println("notfound")
	}

	w.Write([]byte("Hello register"))
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	_ = utils.GetQuery(r, "id")

	ctx := r.Context()
	_, err := h.AuthService.Login(ctx, uuid.New())
	if err != nil {
		log.Println("notfound")
	}

	w.Write([]byte("Hello Login"))
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Logout"))
}
