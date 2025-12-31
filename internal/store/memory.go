// memory.go
//
// Author: Joseph Egan
// Created: 2025-12-22
// Description: Implements an in-memory JobStore that satisfies the interface
// by storing and retrieving jobs using a map with concurrency protection.
//

package store

import (
	"errors"
	"sync"

	"github.com/MachineCreation/go_pipeline/internal/jobs"
)

type MemoryStore struct {
	mu		sync.Mutex
	jobs map[string]*jobs.Job
}

func NewMemoryStore() *MemoryStore {
	return  &MemoryStore{
		jobs: make(map[string]*jobs.Job),
	}
}

func (memoryStore *MemoryStore) Create(job *jobs.Job) error {
	memoryStore.mu.Lock()
	defer memoryStore.mu.Unlock()

	if job == nil {
		return errors.New("job is nil")
	}

	job.Status = "created"
	memoryStore.jobs[job.ID] = job
	return nil
}

func (memoryStore *MemoryStore) Get(id string) (*jobs.Job, error) {
	memoryStore.mu.Lock()
	defer memoryStore.mu.Unlock()

	job, ok := memoryStore.jobs[id]

	if !ok {
		return nil, errors.New("job not found")
	}

	return job, nil
}

func (memoryStore *MemoryStore) StatusUpdate(job *jobs.Job, status string) (*jobs.Job, error) {
	memoryStore.mu.Lock()
	defer memoryStore.mu.Unlock()

	// verify status not empty
	if status == "" {
		return job, errors.New("new status can not be empty")
	}

	// change status
	job.Status = status

	return job, nil
}