// Package main demonstrates how to use the reclameaqui client library.
package main

import (
	"context"
	"fmt"
	"log"

	reclameaqui "github.com/elizandrodantas/reclame-aqui"
	"github.com/elizandrodantas/reclame-aqui/types"
)

func main() {
	// Create a new client with custom timeout
	client := reclameaqui.New(&types.ReclameAquiOptions{
		TimeoutSeconds: 30,
	})
	// Always close the client to free resources
	defer client.Close()

	// Example 1: Search for companies
	fmt.Println("=== Searching Companies ===")
	searchExample(client)

	// Example 2: Get company details
	fmt.Println("\n=== Company Details ===")
	companyDetailsExample(client)

	// Example 3: List complaints
	fmt.Println("\n=== Listing Complaints ===")
	listComplaintsExample(client)

	// Example 4: Get complaint details
	fmt.Println("\n=== Complaint Details ===")
	complaintDetailsExample(client)
}

// searchExample demonstrates how to search for companies
func searchExample(client *reclameaqui.ReclameAqui) {
	ctx := context.Background()

	// Search with pagination
	results, err := client.SearchCompany(ctx, "Facebook", &types.SearchCompanyOptions{
		Page: 0,
		Size: 5,
	})
	if err != nil {
		log.Printf("Error searching companies: %v", err)
		return
	}

	fmt.Printf("Found %d companies\n", results.Count())

	if !results.IsEmpty() {
		fmt.Println("\nShowing first result:")
		if first := results.GetFirst(); first != nil {
			fmt.Printf("Name: %s\n", first.Name)
			fmt.Printf("Shortname: %s\n", first.Shortname)
			if first.StatusDescription != "" {
				fmt.Printf("Status: %s\n", first.StatusDescription)
			}
			if first.FinalScore != "" {
				fmt.Printf("Score: %s\n", first.FinalScore)
			}
			if first.SolvedPercentual != "" {
				fmt.Printf("Resolved: %s%%\n", first.SolvedPercentual)
			}
			fmt.Printf("Verified: %v\n", first.Verified)
			fmt.Printf("Complaints: %d\n", first.Count)
		}
	}
}

// companyDetailsExample demonstrates how to get detailed company information
func companyDetailsExample(client *reclameaqui.ReclameAqui) {
	ctx := context.Background()

	// Get details using company shortname
	company, err := client.DescribeCompany(ctx, "google-brasil")
	if err != nil {
		log.Printf("Error getting company details: %v", err)
		return
	}

	if company.Name != "" {
		fmt.Printf("Company: %s\n", company.Name)
	}
	if company.Shortname != "" {
		fmt.Printf("Shortname: %s\n", company.Shortname)
	}
	if company.Cnpj != "" {
		fmt.Printf("CNPJ: %s\n", company.Cnpj)
	}
	fmt.Printf("Verified: %v\n", company.Verified)
	if company.SiteURL != "" {
		fmt.Printf("Website: %s\n", company.SiteURL)
	}

	// Display reputation information
	if len(company.Reputation) > 0 {
		rep := company.Reputation[0]
		fmt.Printf("\nReputation:\n")
		if rep.StatusDescription != "" {
			fmt.Printf("  Status: %s\n", rep.StatusDescription)
		}
		if rep.FinalScore != "" {
			fmt.Printf("  Score: %s\n", rep.FinalScore)
		}
		if rep.Complaints != "" {
			fmt.Printf("  Complaints: %s\n", rep.Complaints)
		}
		if rep.AnsweredRate != "" {
			fmt.Printf("  Answered: %s%%\n", rep.AnsweredRate)
		}
		if rep.SolvedRate != "" {
			fmt.Printf("  Solved: %s%%\n", rep.SolvedRate)
		}
	}
}

// listComplaintsExample demonstrates how to list complaints with filters
func listComplaintsExample(client *reclameaqui.ReclameAqui) {
	ctx := context.Background()

	// List complaints with filters
	complaints, err := client.ListComplaints(ctx, &types.ListComplaintsOptions{
		Page:   0,
		Size:   5,
		Status: "ANSWERED", // Filter by answered complaints
	})
	if err != nil {
		log.Printf("Error listing complaints: %v", err)
		return
	}

	fmt.Printf("Found %d complaints\n", complaints.Count())

	if len(complaints.Complaints) > 0 {
		fmt.Println("\nShowing first complaint:")
		complaint := complaints.Complaints[0]
		if complaint.Title != "" {
			fmt.Printf("Title: %s\n", complaint.Title)
		}
		if complaint.Company.Name != "" {
			fmt.Printf("Company: %s\n", complaint.Company.Name)
		}
		if complaint.Status != "" {
			fmt.Printf("Status: %s\n", complaint.Status)
		}
		fmt.Printf("Solved: %v\n", complaint.Solved)
		if complaint.City != "" && complaint.State != "" {
			fmt.Printf("Location: %s, %s\n", complaint.City, complaint.State)
		}
	}

	// Display available categories
	if len(complaints.Categories) > 0 {
		fmt.Printf("\nTotal categories: %d\n", len(complaints.Categories))
		if len(complaints.Categories) > 0 && complaints.Categories[0].Description != "" {
			fmt.Printf("Top category: %s (%d complaints)\n",
				complaints.Categories[0].Description,
				complaints.Categories[0].Count)
		}
	}
}

// complaintDetailsExample demonstrates how to get detailed complaint information
func complaintDetailsExample(client *reclameaqui.ReclameAqui) {
	ctx := context.Background()

	// Get complaint details
	complaint, err := client.DescribeComplaint(ctx, "t0A_VuWhniUmsYBT")
	if err != nil {
		log.Printf("Error getting complaint details: %v", err)
		return
	}

	if complaint.Title != "" {
		fmt.Printf("Title: %s\n", complaint.Title)
	}
	if complaint.Description != "" {
		fmt.Printf("Description: %s\n", complaint.Description)
	}
	if complaint.Status != "" {
		fmt.Printf("Status: %s\n", complaint.Status)
	}
	fmt.Printf("Solved: %v\n", complaint.Solved)
	if complaint.City != "" && complaint.State != "" {
		fmt.Printf("Location: %s, %s\n", complaint.City, complaint.State)
	}
}
