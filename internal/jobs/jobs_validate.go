// validate_jobs.go
//
// Author: Joseph Egan
// Created: 2025-12-22
// Description: validates new jobs to prevent invalid job structures by accident
//

package jobs

import (
	"errors"
)

func (job *Job) Validate() error {
	if job.ID == "" {
		return errors.New("id is required")
	}
	return nil
}