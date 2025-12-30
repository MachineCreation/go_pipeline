// store.go
//
// Author: Joseph Egan
// Created: 2025-12-22
// Description: Simple interface behavior export
//

package store

import (
	"github.com/MachineCreation/go_pipeline/internal/jobs"
)

type JobStore interface {
	Create(job *jobs.Job) error
	Get(id string) (*jobs.Job, error)
}