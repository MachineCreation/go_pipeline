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
func NewJobsCreateHandler(jobStore store.JobStore) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(responseWriter, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var job jobs.Job
		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&job); err != nil {
			http.Error(responseWriter, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := job.Validate(); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}

		if err := jobStore.Create(&job); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}

		responseWriter.WriteHeader(http.StatusCreated)
	}
}