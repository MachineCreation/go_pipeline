// jobs.go
//
// Author: Joseph Egan
// Created: 2025-12-30
// Description: job handling
//

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MachineCreation/go_pipeline/internal/jobs"
	"github.com/MachineCreation/go_pipeline/internal/store"
)

// NewJobsCreateHandler returns an HTTP handler function that creates
// jobs using the provided JobStore.
func JobsCreateHandler(jobStore store.JobStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		
		var job jobs.Job
		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&job); err != nil {
			http.Error(writer, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := job.Validate(); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		if err := jobStore.Create(&job); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)
	}
}