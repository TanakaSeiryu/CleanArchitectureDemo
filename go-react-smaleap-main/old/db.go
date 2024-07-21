package main

import (
	"database/sql"

	"github.com/SeiryuTanaka/go-react-smallapp/internal/database"
)

var Queries *database.Queries

func Connect() error {
	dsn := "host=go-react-smallapp-db-1 port=5432 user=user password=password dbname=go-react-small-app sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	Queries = database.New(db)
	return err
}
