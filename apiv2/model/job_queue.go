// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobQueue the job queue info
//
// swagger:model JobQueue
type JobQueue struct {

	// The count of jobs in the job queue
	Count int64 `json:"count,omitempty"`

	// The type of the job queue
	JobType string `json:"job_type,omitempty"`

	// The latency the job queue (seconds)
	Latency int64 `json:"latency,omitempty"`

	// The paused status of the job queue
	Paused bool `json:"paused"`
}

// Validate validates this job queue
func (m *JobQueue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobQueue) UnmarshalBinary(b []byte) error {
	var res JobQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
