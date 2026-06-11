package main

import (
	"fmt"
	"log"
	"net/http"

	"crud-api/database"
	"crud-api/routes"
)

func main() {

	database.ConnectDB()

	routes.SetupRoutes()

	fmt.Println("🚀 Server running on http://localhost:8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}