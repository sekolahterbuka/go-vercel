package handler

import (
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	Interface "github.com/sekolahkita/go-api/server/interface"
	pg "github.com/sekolahkita/go-api/server/repository/postgres/sqlc"
	"github.com/sekolahkita/go-api/server/serializer/json"
	"github.com/sekolahkita/go-api/server/util"
	apperrors "github.com/sekolahkita/go-api/server/util/apperror"
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

type RegisterParams struct {
	Username string `json:"username" validate:"required,min=2,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

// REGISTER
func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		apperrors.NewInternal(w)
		return
	}

	//body parser
	payload, err := json.Decode[RegisterParams](reqBody)
	if err != nil {
		apperrors.NewBadRequest(w, "invalid Input")
		return
	}

	//Validator
	validate := validator.New()
	errValidate := validate.Struct(payload)
	if errValidate != nil {
		apperrors.NewBadRequest(w, errValidate.Error())
		return
	}

	//Register Service
	result, err := h.AuthService.Register(ctx, pg.RegisterParams{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		apperrors.NewConflict(w, "value", payload.Username)
		return
	}

	//Encode Json Response
	resJson, err := json.Encode(&result)
	if err != nil {
		apperrors.NewInternal(w)
		return
	}

	util.SetupResponse(w, []byte(resJson), http.StatusCreated)
}

// LOGIN
func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paramsId := util.GetQuery(r, "id")

	//Login Service
	res, err := h.AuthService.Login(ctx, uuid.MustParse(paramsId))
	if err != nil {
		apperrors.NewNotFound(w, "identifier", paramsId)
		return
	}

	//Encode Json Response
	resJson, err := json.Encode(&res)
	if err != nil {
		apperrors.NewInternal(w)
		return
	}

	util.SetupResponse(w, []byte(resJson), http.StatusOK)
}

// LOGOUT
func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(w, []byte("success"), http.StatusOK)
}
