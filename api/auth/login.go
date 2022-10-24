package handler

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sekolahkita/go-api/server/handler"
	"github.com/sekolahkita/go-api/server/repository"
	"github.com/sekolahkita/go-api/server/service"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", repository.PgConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	service := service.NewAuthService(&service.AuthConfig{Querier: repo})
	handler := handler.NewAuthHandler(&handler.AuthConfig{
		AuthService: service,
	})

	switch r.Method {

	case "POST":
		handler.Login(w, r)

	default:
		http.Error(w, "method "+r.Method+" not allowed", http.StatusBadRequest)
	}

}
