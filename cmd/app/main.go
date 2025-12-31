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

	"github.com/MachineCreation/go_pipeline/internal/store"
	"github.com/MachineCreation/go_pipeline/internal/handlers"
	"github.com/MachineCreation/go_pipeline/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	//-----------------------------------------------------------------------------
	//Resources
	//-----------------------------------------------------------------------------
	
	// create new memory store
	jobStore := store.NewMemoryStore()

	// create new router
	router := chi.NewRouter()

	// middleware for logging requests
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	//-----------------------------------------------------------------------------
	//Handlers
	//-----------------------------------------------------------------------------

	//standard health check to determine if server is alive
	router.Get("/health", handlers.HealthHandler)

	//create new job
	router.Post("/jobs", handlers.JobsCreateHandler(jobStore))

	//Get job by id
	router.Get("/jobs/{id}", handlers.JobsGetByIDHandler(jobStore))

	//cancel job by id
	router.Put("/jobs/{id}/cancel", handlers.JobsCancelByIDHandler(jobStore))

	//-----------------------------------------------------------------------------
	//Server
	//-----------------------------------------------------------------------------

	serverError := server.StartHTTPServer(":8080", router)
	if serverError != nil {
		log.Fatal(serverError)
	}
}