package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)

func main() {
	var err error

	db, err := sql.Open("postgres", "port=5433 user=postgres password=postgres dbname=bookstore sslmode=disable")
	if err != nil {
		fmt.Println("sql.Open err ")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("db.Ping err ")
	}

	fmt.Println("SUCC")
	// ...
}
