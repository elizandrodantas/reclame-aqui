package reclameaqui

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elizandrodantas/reclame-aqui/types"
)

// SearchCompany searches for companies by name.
// It returns a list of companies matching the search criteria.
//
// Parameters:
//   - ctx: Context for request cancellation and timeout
//   - name: Company name to search for (required)
//   - opt: Search options including pagination (can be nil for defaults)
//
// Returns the search results or an error if the request fails.
func (r *ReclameAqui) SearchCompany(ctx context.Context, name string, opt *types.SearchCompanyOptions) (*types.SearchCompanyResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if name == "" {
		return nil, ErrSearchCompanyNameEmpty
	}

	// Apply defaults if options not provided
	if opt == nil {
		opt = &types.SearchCompanyOptions{
			Page: 0,
			Size: defaultSize,
		}
	}

	// Validate options
	if opt.Page < 0 {
		return nil, ErrPageInvalid
	}

	if opt.Size <= 0 {
		opt.Size = defaultSize
	}

	if opt.Size > maxSize {
		opt.Size = maxSize
	}

	params := map[string]string{
		"text": name,
		"page": fmt.Sprintf("%d", opt.Page),
		"size": fmt.Sprintf("%d", opt.Size),
	}

	fullURL, err := r.client.BuildURL(endpointSearchCompany, params)
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

	var response types.SearchCompanyResponse
	if err := json.Unmarshal(res.BodyBytes, &response); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalJSON, err)
	}

	return &response, nil
}
