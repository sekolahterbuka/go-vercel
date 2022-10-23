package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/sekolahkita/go-api/server/model"
	"github.com/sekolahkita/go-api/server/serializer/json"
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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error")
	}

	data, err := json.Decode[model.RegisterParams](reqBody)
	if err != nil {
		log.Println("error")
	}

	ctx := r.Context()
	result, err := h.AuthService.Register(ctx, model.RegisterParams{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		log.Println("notfound")
	}
	log.Printf("%v save to database", result)

	resJson, err := json.Encode[model.Auth](&model.Auth{
		UID:      uuid.New(),
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		log.Println("notfound")
	}

	w.Write([]byte(resJson))
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	_ = utils.GetQuery(r, "id")

	ctx := r.Context()
	_, err := h.AuthService.Login(ctx, uuid.New())
	if err != nil {
		log.Println("notfound")
	}

	resJson, err := json.Encode[model.Auth](&model.Auth{
		UID:      uuid.New(),
		Username: "sanja",
		Email:    "sanja@mail.com",
		Password: "pasword",
	})
	if err != nil {
		log.Println("notfound")
	}

	w.Write([]byte(resJson))
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Logout"))
}
