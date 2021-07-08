package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"yogo.io/go-tripPlanner-backend/api"
	yogo "yogo.io/go-tripPlanner-backend/db"
	"yogo.io/go-tripPlanner-backend/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := yogo.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
