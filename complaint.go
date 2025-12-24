package reclameaqui

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elizandrodantas/reclame-aqui/types"
)

// ListComplaints retrieves a list of complaints with optional filters.
//
// Parameters:
//   - ctx: Context for request cancellation and timeout
//   - opt: Filtering and pagination options (can be nil for defaults)
//
// Returns a list of complaints matching the filters or an error if the request fails.
func (r *ReclameAqui) ListComplaints(ctx context.Context, opt *types.ListComplaintsOptions) (*types.ListComplaintsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Apply defaults if options not provided
	if opt == nil {
		opt = &types.ListComplaintsOptions{
			Page: 0,
			Size: defaultPerPage,
		}
	}

	// Validate options
	if opt.Page < 0 {
		return nil, ErrPageInvalid
	}

	if opt.Size <= 0 {
		opt.Size = defaultPerPage
	}

	if opt.Size > maxPerPage {
		opt.Size = maxPerPage
	}

	params := map[string]string{
		"page":    fmt.Sprintf("%d", opt.Page),
		"perPage": fmt.Sprintf("%d", opt.Size),
	}

	if opt.CompanyID != 0 {
		params["companyId"] = fmt.Sprintf("%d", opt.CompanyID)
	}

	if opt.Status != "" {
		params["status"] = opt.Status
	}

	if opt.Category != "" {
		params["category"] = opt.Category
	}

	if opt.ProblemTypeID != "" {
		params["problemType"] = opt.ProblemTypeID
	}

	fullURL, err := r.client.BuildURL(endpointListComplaints, params)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateRequest, err)
	}

	res, err := r.client.DoRequest(fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedRequest, err)
	}

	if res.Status != 200 {
		return nil, r.client.parseAPIError(res)
	}

	var response types.ListComplaintsResponse
	if err := json.Unmarshal(res.BodyBytes, &response); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalJSON, err)
	}

	return &response, nil
}
