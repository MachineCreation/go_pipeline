// main.go
//
// Author: Joseph Egan
// Created: 2025-12-22
// Description: Acts as the application entry point, wiring together
// dependencies and starting the program’s execution flow.
//

package main

import (
	"errors"
	"fmt"

	"github.com/MachineCreation/go_pipeline/internal/jobs"
)

func NewJob(id string) (*jobs.Job, error) {

	// handle empty id string
	if id == "" {
		return nil, errors.New("id is required")
	}

	// no errors, return default value pair
	return &jobs.Job{
		ID:		id,
		Status: "Queued",
	}, nil
}

func main() {
	job, err := NewJob("job-001")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created job: %+v\n", *job)
}