package reclameaqui

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elizandrodantas/reclame-aqui/types"
)

// DescribeCompany retrieves detailed information about a specific company.
//
// Parameters:
//   - ctx: Context for request cancellation and timeout
//   - slug: Company slug identifier (required)
//
// Returns detailed company information or an error if the request fails.
func (r *ReclameAqui) DescribeCompany(ctx context.Context, slug string) (*types.DescribeCompanyResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if slug == "" {
		return nil, ErrDescribeCompanySlugEmpty
	}

	fullURL, err := r.client.BuildURL(fmt.Sprintf(endpointDescribeCompany, slug), nil)
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

	var response types.DescribeCompanyResponse
	if err := json.Unmarshal(res.BodyBytes, &response); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalJSON, err)
	}

	return &response, nil
}
