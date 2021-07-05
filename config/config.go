package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"yogo.io/go-tripPlanner-backend/helpers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Sc0rpion"
	dbname   = "yogo"
)

func GetDB() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)

	helpers.HandleErr(err)
	return
}
