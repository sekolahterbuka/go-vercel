package repository

import (
	"database/sql"
	"fmt"

	pg "github.com/sekolahkita/go-api/server/repository/postgres/sqlc"
)

const (
	host     = "containers-us-west-94.railway.app"
	port     = 6643
	user     = "postgres"
	password = "2q4qsL4Qd6g41UfrLMZ1"
	dbname   = "railway"
)

var PgConn = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

type repository struct {
	*pg.Queries
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		DB:      db,
		Queries: pg.New(db),
	}
}
