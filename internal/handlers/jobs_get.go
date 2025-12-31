// internal\handlers\getJobs.go
//
// Author: Joseph Egan
// Created: 2025-12-31
// Description: functions for searching existing jobs
//

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MachineCreation/go_pipeline/internal/store"
	"github.com/go-chi/chi/v5"
)

// JobsGetByID returns an http handler function that gets a jobs
// details from the provided JobStore using its provided id
func JobsGetByIDHandler(jobStore store.JobStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jobID := chi.URLParam(request, "id")

		// check for id in request
		if jobID == "" {
			http.Error(writer, "request missing job id", http.StatusBadRequest)
			return
		}

		// search for job in JobStore
		job, err := jobStore.Get(jobID)

		// handle error
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}

		// handle found jobID
		// write response header to "OK"
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		// encode response
		encoder := json.NewEncoder(writer)

		// handle encoding error
		if err := encoder.Encode(job); err != nil {
			http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}