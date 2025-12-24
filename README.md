<div align="center">
  <img src="https://storage.googleapis.com/reclameaqui-assets/images/logo-25-anos.svg" alt="Reclame Aqui Logo" width="200"/>
  
  # Reclame Aqui Go Client

  [![Go Reference](https://pkg.go.dev/badge/github.com/elizandrodantas/reclame-aqui.svg)](https://pkg.go.dev/github.com/elizandrodantas/reclame-aqui)
  [![Go Report Card](https://goreportcard.com/badge/github.com/elizandrodantas/reclame-aqui)](https://goreportcard.com/report/github.com/elizandrodantas/reclame-aqui)
</div>

A Go client library for the Reclame AQUI API.

## Features

- ✅ **Type-safe**: Strongly typed Go structures with full documentation
- ✅ **Easy to use**: Clean and intuitive API
- ✅ **Configurable**: Customizable timeout and request options
- ✅ **Well-documented**: Complete GoDoc documentation

## Installation

```bash
go get github.com/elizandrodantas/reclame-aqui
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    reclameaqui "github.com/elizandrodantas/reclame-aqui"
    "github.com/elizandrodantas/reclame-aqui/types"
)

func main() {
    // Create client with custom timeout
    client := reclameaqui.New(&types.ReclameAquiOptions{
        TimeoutSeconds: 30,
    })
    defer client.Close()

    // Search for companies
    results, err := client.SearchCompany(context.Background(), "Google Brasil", &types.SearchCompanyOptions{
        Page: 0,
        Size: 10,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Display results
    for _, company := range *results {
        fmt.Printf("%s - %s (Score: %s)\n", 
            company.Name, 
            company.StatusDescription, 
            company.FinalScore)
    }
}
```

## API Reference

### Client Creation

```go
// Create a new client
client := reclameaqui.New(&types.ReclameAquiOptions{
    TimeoutSeconds: 30, // Optional, defaults to 30 seconds
})
defer client.Close() // Always close when done
```

### Search Companies

```go
results, err := client.SearchCompany(ctx, "company name", &types.SearchCompanyOptions{
    Page: 0,    // Page number (starts at 0)
    Size: 10,   // Results per page (max 50)
})
```

### Get Company Details

```go
company, err := client.DescribeCompany(ctx, "company-slug")
```

### List Complaints

```go
complaints, err := client.ListComplaints(ctx, &types.ListComplaintsOptions{
    CompanyID:     12345,      // Optional: filter by company ID
    Status:        "ANSWERED",  // Optional: ANSWERED, PENDING, EVALUATED, SOLVED
    Page:          0,           // Page number
    Size:          10,          // Results per page (max 10)
    Category:      "",          // Optional: category ID
    ProblemTypeID: "",          // Optional: problem type ID
})
```

### Get Complaint Details

```go
complaint, err := client.DescribeComplaint(ctx, "complaint-id")
```

## Data Structures

### Company

```go
type Company struct {
    ID                   string
    Name                 string
    Shortname            string
    Status               string   // GREAT, GOOD, NO_INDEX, NOT_RECOMMENDED
    StatusDescription    string
    FinalScore           string
    SolvedPercentual     string
    Logo                 string
    City                 string
    State                string
    Reputation           string
    URL                  *string
    Verified             bool
    Count                int      // Number of complaints
}
```

### Complaint

```go
type Complaint struct {
    ID          string
    Title       string
    Description string
    Solved      bool
    City        string
    State       string
    Status      string
    Created     string
    Company     CompanyReference
}
```

## Error Handling

The library provides sentinel errors for better error handling:

```go
import "errors"

results, err := client.SearchCompany(ctx, "", nil)
if err != nil {
    if errors.Is(err, reclameaqui.ErrSearchCompanyNameEmpty) {
        // Handle empty name error
    }
    if errors.Is(err, reclameaqui.ErrFailedRequest) {
        // Handle request failure
    }
}
```

Available errors:
- `ErrCreateRequest` - HTTP request creation failed
- `ErrFailedRequest` - HTTP request execution failed
- `ErrUnmarshalJSON` - JSON parsing failed
- `ErrClientError` - API returned an error
- `ErrPageInvalid` - Invalid page number
- `ErrSearchCompanyNameEmpty` - Empty company name
- `ErrDescribeCompanySlugEmpty` - Empty company slug
- `ErrDescribeComplaintIDInvalid` - Invalid complaint ID

## Best Practices

### Always Close the Client

```go
client := reclameaqui.New(nil)
defer client.Close() // Release resources
```

### Use Context for Cancellation

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

results, err := client.SearchCompany(ctx, "company", nil)
```

### Handle Pagination

```go
page := 0
for {
    results, err := client.SearchCompany(ctx, "company", &types.SearchCompanyOptions{
        Page: page,
        Size: 50,
    })
    if err != nil {
        break
    }
    if results.IsEmpty() {
        break
    }
    // Process results
    page++
}
```

## Project Structure

```
.
├── client.go              # HTTP client implementation
├── constants.go           # Package constants
├── error.go              # Error definitions
├── reclame-aqui.go       # Main client interface
├── company.go            # Company search methods
├── company_describe.go   # Company details methods
├── complaint.go          # Complaint listing methods
├── complaint_describe.go # Complaint details methods
└── domain/               # Data models
    ├── describe.go       # Detailed response models
    ├── list.go          # List response models
    ├── options.go       # Request options
    ├── search.go        # Search response models
    ├── shared.go        # Shared types
    └── status.go        # Status types
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This is an unofficial client library for the Reclame AQUI API. It is not affiliated with, endorsed by, or in any way associated with Reclame AQUI.
