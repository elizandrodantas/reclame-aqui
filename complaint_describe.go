package reclameaqui

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elizandrodantas/reclame-aqui/types"
)

// DescribeComplaint retrieves detailed information about a specific complaint.
//
// Parameters:
//   - ctx: Context for request cancellation and timeout
//   - id: Complaint ID (required)
//
// Returns detailed complaint information or an error if the request fails.
func (r *ReclameAqui) DescribeComplaint(ctx context.Context, id string) (*types.DescribeComplaintResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if id == "" {
		return nil, ErrDescribeComplaintIDInvalid
	}

	fullURL, err := r.client.BuildURL(fmt.Sprintf(endpointDescribeComplaint, id), nil)
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

	var response types.DescribeComplaintResponse
	if err := json.Unmarshal(res.BodyBytes, &response); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalJSON, err)
	}

	return &response, nil
}
