package main

import (
	"assyifa/config"
	"assyifa/master"
)

func main() {
	db, _, dbHost, portServer := config.Connection()
	router := config.CreateRouter()

	master.Init(router, db)
	config.RunServer(router, dbHost, portServer)
}
