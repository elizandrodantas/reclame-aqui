package reclameaqui

import "errors"

// Sentinel errors for better error handling and testing
var (
	// ErrCreateRequest is returned when HTTP request creation fails
	ErrCreateRequest = errors.New("failed to create HTTP request")

	// ErrFailedRequest is returned when HTTP request execution fails
	ErrFailedRequest = errors.New("HTTP request failed")

	// ErrUnmarshalJSON is returned when JSON unmarshaling fails
	ErrUnmarshalJSON = errors.New("failed to unmarshal JSON response")

	// ErrClientError is returned when API returns an error response
	ErrClientError = errors.New("API client error")

	// ErrPageInvalid is returned when page number is negative
	ErrPageInvalid = errors.New("page number cannot be negative")

	// ErrSearchCompanyNameEmpty is returned when company name is empty
	ErrSearchCompanyNameEmpty = errors.New("company name cannot be empty")

	// ErrDescribeCompanySlugEmpty is returned when company slug is empty
	ErrDescribeCompanySlugEmpty = errors.New("company slug cannot be empty")

	// ErrDescribeComplaintIDInvalid is returned when complaint ID is invalid
	ErrDescribeComplaintIDInvalid = errors.New("complaint ID is required")
)
