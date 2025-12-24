// Package reclameaqui provides a Go client for the Reclame AQUI API.
//
// Reclame AQUI is a Brazilian consumer complaint platform where users can
// report issues with companies and track their resolution status.
//
// This client library provides type-safe access to the Reclame AQUI API with
// built-in Cloudflare bypass using CycleTLS.
//
// # Basic Usage
//
// Create a client and search for companies:
//
//	client := reclameaqui.New(&types.ReclameAquiOptions{
//	    TimeoutSeconds: 30,
//	})
//	defer client.Close()
//
//	results, err := client.SearchCompany(context.Background(), "Company Name", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, company := range *results {
//	    fmt.Printf("%s - %s\n", company.Name, company.StatusDescription)
//	}
//
// # Features
//
//   - Search for companies by name
//   - Get detailed company information
//   - List complaints with filters
//   - Get detailed complaint information
//   - Built-in Cloudflare bypass
//   - Configurable timeouts
//   - Type-safe error handling
//
// # Error Handling
//
// The package provides sentinel errors for better error handling:
//
//	results, err := client.SearchCompany(ctx, "", nil)
//	if errors.Is(err, reclameaqui.ErrSearchCompanyNameEmpty) {
//	    // Handle empty name error
//	}
//
// # Pagination
//
// All list methods support pagination:
//
//	for page := 0; ; page++ {
//	    results, err := client.SearchCompany(ctx, "name", &types.SearchCompanyOptions{
//	        Page: page,
//	        Size: 50,
//	    })
//	    if err != nil || results.IsEmpty() {
//	        break
//	    }
//	    // Process results
//	}
//
// # Context Support
//
// All methods accept a context for cancellation and timeout:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	results, err := client.SearchCompany(ctx, "name", nil)
//
// # Resource Management
//
// Always close the client when done to free resources:
//
//	client := reclameaqui.New(nil)
//	defer client.Close()
package reclameaqui
