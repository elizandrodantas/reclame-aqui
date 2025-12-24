package reclameaqui

import (
	"encoding/json"
	"fmt"
	"maps"
	"net/url"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/elizandrodantas/reclame-aqui/types"
)

// HTTPClient defines the interface for making HTTP requests
type HTTPClient interface {
	DoRequest(fullURL string, headers map[string]string) (*cycletls.Response, error)
	BuildURL(endpoint string, params map[string]string) (string, error)
	Close()
}

// Client implements the HTTPClient interface using CycleTLS
type Client struct {
	cycleTLS cycletls.CycleTLS
	baseURL  string
}

// newClient creates and returns a new instance of the Reclame AQUI API client.
// This function is package-private and should be called through New().
func newClient(timeout int64) *Client {
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	client := cycletls.Init()

	return &Client{
		cycleTLS: client,
		baseURL:  baseURL,
	}
}

// DoRequest performs an HTTP GET request using CycleTLS and returns the response.
// It automatically adds default headers to mimic a real browser.
func (c *Client) DoRequest(fullURL string, headers map[string]string) (*cycletls.Response, error) {
	defaultHeaders := map[string]string{
		"User-Agent": userAgent,
	}

	// Merge custom headers with defaults
	maps.Copy(defaultHeaders, headers)

	response, err := c.cycleTLS.Do(fullURL, cycletls.Options{
		Body:      "",
		Ja3:       ja3Fingerprint,
		UserAgent: defaultHeaders["User-Agent"],
		Headers:   defaultHeaders,
	}, "GET")

	if err != nil {
		return nil, fmt.Errorf("cycletls request failed: %w", err)
	}

	return &response, nil
}

// BuildURL constructs the full URL with the given endpoint and query parameters.
func (c *Client) BuildURL(endpoint string, params map[string]string) (string, error) {
	fullURL := c.baseURL + endpoint

	if len(params) > 0 {
		urlObj, err := url.Parse(fullURL)
		if err != nil {
			return "", fmt.Errorf("failed to parse URL: %w", err)
		}

		q := urlObj.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		urlObj.RawQuery = q.Encode()
		fullURL = urlObj.String()
	}

	return fullURL, nil
}

// parseAPIError attempts to parse error responses from the API
func (c *Client) parseAPIError(res *cycletls.Response) error {
	var errException types.ErrorClientExceptionResponse
	if err := json.Unmarshal(res.BodyBytes, &errException); err == nil {
		return fmt.Errorf("%w: %s", ErrClientError, errException.Message)
	}

	var errSimple types.ErrorClientSimpleResponse
	if err := json.Unmarshal(res.BodyBytes, &errSimple); err == nil {
		return fmt.Errorf("%w: %s", ErrClientError, errSimple.Message)
	}

	return fmt.Errorf("%w: status code %d", ErrClientError, res.Status)
}

// Close closes the CycleTLS client to free resources.
// Should be called when the client is no longer needed.
func (c *Client) Close() {
	c.cycleTLS.Close()
}
