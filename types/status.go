package types

import "errors"

// ComplaintStatus represents the status of a complaint.
type ComplaintStatus string

// ErrInvalidComplaintStatus is returned when an invalid complaint status is provided.
var ErrInvalidComplaintStatus = errors.New("invalid complaint status")

// Complaint status constants
const (
	// StatusComplaintAnswered indicates the complaint has been answered by the company
	StatusComplaintAnswered ComplaintStatus = "ANSWERED"

	// StatusComplaintPending indicates the complaint is awaiting response
	StatusComplaintPending ComplaintStatus = "PENDING"

	// StatusComplaintEvaluated indicates the complaint has been evaluated
	StatusComplaintEvaluated ComplaintStatus = "EVALUATED"

	// StatusComplaintSolved indicates the complaint has been resolved
	StatusComplaintSolved ComplaintStatus = "SOLVED"
)

// IsValid checks if the complaint status is valid.
func (s ComplaintStatus) IsValid() bool {
	switch s {
	case StatusComplaintAnswered, StatusComplaintPending, StatusComplaintEvaluated, StatusComplaintSolved:
		return true
	default:
		return false
	}
}

// String returns the string representation of the complaint status.
func (s ComplaintStatus) String() string {
	return string(s)
}

// NewComplaintStatus creates a new ComplaintStatus from a string.
// Returns an error if the status is invalid.
func NewComplaintStatus(status string) (ComplaintStatus, error) {
	cs := ComplaintStatus(status)
	if !cs.IsValid() {
		return "", ErrInvalidComplaintStatus
	}

	return cs, nil
}
