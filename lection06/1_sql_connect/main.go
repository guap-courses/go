package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://username:password@localhost:5442/guap_course?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

}
