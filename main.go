package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/luisteixeira/simplebank/api"
	db "github.com/luisteixeira/simplebank/db/sqlc"
	"github.com/luisteixeira/simplebank/util"
)

func main() {
	config, err := util.LoadCofig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}