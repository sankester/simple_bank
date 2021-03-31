package main

import (
	"backend-master-class/api"
	db "backend-master-class/db/sqlc"
	"backend-master-class/utils"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cnnot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect db : ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect  server", err)
	}
}
