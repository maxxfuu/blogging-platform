package main

import (
	"log"

	"blogging-platform/api"
	"blogging-platform/database"
)

func main() {

	database.InitDB()

	router := api.SetupRouter()
	log.Println("Starting server on :8080...")
	router.Run(":8080")
}
