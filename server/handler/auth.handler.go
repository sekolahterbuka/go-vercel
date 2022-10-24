package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	Interface "github.com/sekolahkita/go-api/server/interface"
	pg "github.com/sekolahkita/go-api/server/repository/postgres/sqlc"
	"github.com/sekolahkita/go-api/server/serializer/json"
	"github.com/sekolahkita/go-api/server/utils"
)

type authHandler struct {
	AuthService Interface.AuthService
}
type AuthConfig struct {
	AuthService Interface.AuthService
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

// REGISTER
func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Decode[pg.RegisterParams](reqBody)
	if err != nil {
		log.Println("error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.AuthService.Register(ctx, pg.RegisterParams{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		log.Println("notfound")
		w.WriteHeader(http.StatusConflict)
		return
	}
	log.Printf("%v save to database", result)

	resJson, err := json.Encode[pg.Auth](&result)
	if err != nil {
		log.Println("notfound")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(resJson))
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paramsId := utils.GetQuery(r, "id")

	res, err := h.AuthService.Login(ctx, uuid.MustParse(paramsId))
	if err != nil {
		log.Println("notfound")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resJson, err := json.Encode[pg.Auth](&res)
	if err != nil {
		log.Println("notfound")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(resJson))
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Logout"))
}
