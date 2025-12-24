package types

// ReclameAquiOptions contains configuration options for the Reclame AQUI client.
type ReclameAquiOptions struct {
	// TimeoutSeconds is the HTTP request timeout in seconds.
	// If 0 or negative, the default timeout (30 seconds) will be used.
	TimeoutSeconds int
}

// SearchCompanyOptions contains options for searching companies.
type SearchCompanyOptions struct {
	// Page is the page number for pagination (starting from 0).
	Page int

	// Size is the number of results per page.
	// Maximum value is 50. If 0 or negative, defaults to 10.
	Size int
}

// ListComplaintsOptions contains filtering and pagination options for listing complaints.
type ListComplaintsOptions struct {
	// CompanyID filters complaints by company.
	// If 0, lists complaints from all companies.
	// This is the numeric ID of the company.
	CompanyID int

	// Status filters complaints by status.
	// Possible values: ANSWERED, PENDING, EVALUATED, SOLVED
	Status string

	// Page is the page number for pagination (starting from 0).
	Page int

	// Category is the category ID to filter complaints.
	Category string

	// ProblemTypeID is the problem type ID to filter complaints.
	ProblemTypeID string

	// Size is the number of results per page.
	// Maximum value is 10. If 0 or negative, defaults to 10.
	Size int
}
