package types

// Company represents a company in the Reclame AQUI platform.
type Company struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Shortname            string   `json:"shortname"`
	Created              string   `json:"created"`
	SolvedPercentual     string   `json:"solvedPercentual"`
	Status               string   `json:"status"`
	StatusDescription    string   `json:"statusDescription"`
	FinalScore           string   `json:"finalScore"`
	Logo                 string   `json:"logo"`
	City                 string   `json:"city"`
	State                string   `json:"state"`
	RavLabel             string   `json:"ravLabel"`
	SegmentShortname     string   `json:"segmentShortname"`
	MainSegmentShortname string   `json:"mainSegmentShortname"`
	Count                int      `json:"count"`
	Reputation           string   `json:"reputation"`
	URL                  *string  `json:"url"`
	Documents            []string `json:"documents"`
	Verified             bool     `json:"verified"`
	HasVerified          bool     `json:"hasVerified"`
	HasNotVerified       bool     `json:"hasNotVerified"`
}

// SearchCompanyResponse is a collection of companies returned from a search query.
type SearchCompanyResponse []Company

// Count returns the number of companies in the search response.
func (s *SearchCompanyResponse) Count() int {
	return len(*s)
}

// IsEmpty checks if the search response contains no companies.
func (s *SearchCompanyResponse) IsEmpty() bool {
	return s.Count() == 0
}

// GetFirst returns the first company in the search response, or nil if empty.
func (s *SearchCompanyResponse) GetFirst() *Company {
	if s.IsEmpty() {
		return nil
	}

	return &(*s)[0]
}
