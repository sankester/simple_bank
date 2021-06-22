package main

import (
	"database/sql"
	"log"

	"github.com/sankester/simple_bank/api"
	db "github.com/sankester/simple_bank/db/sqlc"
	"github.com/sankester/simple_bank/utils"

	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
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
