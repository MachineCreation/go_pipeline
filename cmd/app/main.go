// main.go
//
// Author: Joseph Egan
// Created: 2025-12-22
// Description: Acts as the application entry point, wiring together
// dependencies and starting the program’s execution flow.
//

package main

import (
	"log"
	"net/http"

	"github.com/MachineCreation/go_pipeline/internal/store"
	"github.com/MachineCreation/go_pipeline/internal/handlers"
	"github.com/MachineCreation/go_pipeline/internal/server"
)

func main() {

	//-----------------------------------------------------------------------------
	//Resources
	//-----------------------------------------------------------------------------
	
	// create new memory store
	jobStore := store.NewMemoryStore()

	//-----------------------------------------------------------------------------
	//Handlers
	//-----------------------------------------------------------------------------

	//standard health check to determine if server is alive
	http.HandleFunc("/health", handlers.HealthHandler)

	//create new job
	http.HandleFunc("/jobs", handlers.NewJobsCreateHandler(jobStore))

	//-----------------------------------------------------------------------------
	//Server
	//-----------------------------------------------------------------------------

	serverError := server.StartHTTPServer(":8080", nil)
	if serverError != nil {
		log.Fatal(serverError)
	}
}