package main

import (
	"log"
	"net/http"

	"task/database"
	"task/routes"

)

func main() {

	log.Println("Running the main file ")
	routes.Routes()

	log.Println("Database has been created")
	database.Create_DB()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server started")
}
