package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"yogo.io/go-tripPlanner-backend/api"
	yogo "yogo.io/go-tripPlanner-backend/db"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:{Password}@localhost/yogo?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := yogo.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
