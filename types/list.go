package types

// Complaint represents a single complaint in the list response
type Complaint struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Solved      bool             `json:"solved"`
	City        string           `json:"city"`
	State       string           `json:"state"`
	Status      string           `json:"status"`
	Created     string           `json:"created"`
	Company     CompanyReference `json:"company"`
}

// CompanyReference contains basic company information
type CompanyReference struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
}

// Category represents a complaint category with count
type Category struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

// Product represents a product with complaint count
type Product struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

// Problem represents a problem type with complaint count
type Problem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

// ListComplaintsResponse represents the response from listing complaints
type ListComplaintsResponse struct {
	Complaints []Complaint `json:"complaints"`
	Categories []Category  `json:"categories"`
	Products   []Product   `json:"products"`
	Problems   []Problem   `json:"problems"`
}

// Count returns the number of complaints in the response.
func (r *ListComplaintsResponse) Count() int {
	return len(r.Complaints)
}

// IsEmpty checks if the response contains no complaints.
func (r *ListComplaintsResponse) IsEmpty() bool {
	return r.Count() == 0
}

// GetFirst returns the first complaint in the response, or nil if empty.
func (r *ListComplaintsResponse) GetFirst() *Complaint {
	if r.IsEmpty() {
		return nil
	}

	return &r.Complaints[0]
}
