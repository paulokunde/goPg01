package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"app/api"
	"app/modelo"
)

func main() {
	modelo.LoadDbParameters()
	driver := modelo.Parametros.DbDriver
	source := modelo.Parametros.DbSource
	serverAddress := "0.0.0.0:8000"

	conn, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal("cannot connect db", err)
	}

	store := modelo.ExecuteNewStore(conn)
	server := api.InstanceServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("api iniciou com erro", err)
	}

}
