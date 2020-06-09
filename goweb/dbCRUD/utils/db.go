package utils

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("postgres", "port=5433 user=postgres password=postgres dbname=bookstore sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	}
}
