package main

import (
	"log"
	"net/http"

	"yehtun.com/rest-api-crud/config"
	"yehtun.com/rest-api-crud/routes"
)

func main() {
    // Initialize database connection
    config.ConnectDB()

    // Set up routes
    router := routes.SetupRoutes()

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}