package main

import (
	"log"
	"net/http"

	"github.com/esuEdu/casa-oliveira/internal/config"
	"github.com/esuEdu/casa-oliveira/routes"
)

func main() {
	db := config.InitDB()

	router := http.NewServeMux()

	routes.SetupRoutes(router, db)

	// TODO: change the string liteteral of port 8080 to use env variable
	log.Print("Server listening on http://localhost:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
