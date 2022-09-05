package main

import (
	"cryptovote/database"
	"cryptovote/server"
	"fmt"
	"log"
	"os"
)

func main() {
	os.Setenv("PostgresConn", "host=localhost port=5431 user=postgres password=postgres dbname=cryptovote sslmode=disable")

	if err := database.StartDatabase(os.Getenv("PostgresConn")); err != nil {
		log.Panicln(err)
	}

	fmt.Println("running server")
	server.StartServer("8081")
}
