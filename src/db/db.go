package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=samarino_loja password=87272383 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
