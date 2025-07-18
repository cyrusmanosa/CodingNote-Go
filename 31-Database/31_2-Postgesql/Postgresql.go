package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DBSourceInfo = "postgresql://root:secret@localhost:5432/dating_date_info?sslmode=disable"

func main() {
	// Info
	info_conn, err := pgxpool.New(context.Background(), DBSourceInfo)
	if err != nil {
		log.Fatal().Msg("cannot connect to Info indb")
	}
	defer info_conn.Close()
}
