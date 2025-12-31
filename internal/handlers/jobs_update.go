// internal\handlers\jobs_update.go
//
// Author: Joseph Egan
// Created: 2025-12-31
// Description: functions to modify current jobs
//

package handlers

import (
	"net/http"

	"github.com/MachineCreation/go_pipeline/internal/store"
	"github.com/go-chi/chi/v5"
)

// JobsCancelByIDHandler returns an HTTP handler function that gets
// a job from the provided JobStore with the provided ID and changes
// the job's status to canceled
func JobsCancelByIDHandler (jobStore store.JobStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jobID := chi.URLParam(request, "id")

		// check for id in request (error 400)
		if jobID == "" {
			http.Error(writer, "request missing job id", http.StatusBadRequest)
		}

		// search for job in JobStore
		job, err := jobStore.Get(jobID)

		// handle error (404 not found)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}

		// handle found job
		// change status to "canceled"
		jobStore.StatusUpdate(job, "canceled")

		// check status
		// handle error (error 500 internal **for now**)
		if job.Status != "canceled" {
			http.Error(writer, "job status not updated", http.StatusInternalServerError)
			return
		}

		// write response header
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("job canceled successfully"))
	}
}