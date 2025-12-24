package reclameaqui

import "github.com/elizandrodantas/reclame-aqui/types"

// ReclameAqui is the main client for interacting with the Reclame AQUI API.
// It provides methods for searching companies, listing complaints, and retrieving
// detailed information about companies and complaints.
type ReclameAqui struct {
	client *Client
}

// New creates and returns a new instance of the Reclame AQUI API client.
// If options is nil or TimeoutSeconds is 0, the default timeout will be used.
//
// Example:
//
//	client := reclameaqui.New(&types.ReclameAquiOptions{
//	    TimeoutSeconds: 30,
//	})
//	defer client.Close()
func New(options *types.ReclameAquiOptions) *ReclameAqui {
	timeout := int64(0)
	if options != nil && options.TimeoutSeconds > 0 {
		timeout = int64(options.TimeoutSeconds)
	}

	return &ReclameAqui{
		client: newClient(timeout),
	}
}

// Close closes the underlying HTTP client and releases resources.
// It should be called when the ReclameAqui client is no longer needed.
func (r *ReclameAqui) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
