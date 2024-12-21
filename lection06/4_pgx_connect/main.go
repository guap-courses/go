package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	dsn := "postgres://username:password@localhost:5442/guap_course?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	if err = conn.Ping(context.Background()); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}
}
