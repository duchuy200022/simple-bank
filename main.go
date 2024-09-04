package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/duchuy200022/simplebank/api"
	db "github.com/duchuy200022/simplebank/db/sqlc"
	"github.com/duchuy200022/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Println("DB Driver: ", config.DBDriver)
	log.Println("DB Source: ", config.DBSource)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
