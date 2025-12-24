package types

type DescribeCompanyResponse struct {
	ID          string        `json:"id"`
	LegacyID    int           `json:"legacyId"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Cnpj        string        `json:"cnpj"`
	Status      string        `json:"status"`
	Logo        string        `json:"logo"`
	Verified    bool          `json:"verified"`
	IsRaForms   bool          `json:"isRaForms"`
	SiteURL     string        `json:"siteUrl"`
	Created     string        `json:"created"`
	CoverURL    string        `json:"coverUrl"`
	VideosURL   []interface{} `json:"videosUrl"`
	Reputation  []struct {
		Main              bool   `json:"main"`
		Type              string `json:"type"`
		Complaints        string `json:"complaints"`
		Answers           string `json:"answers"`
		Ratings           string `json:"ratings"`
		AnsweredRate      string `json:"answeredRate"`
		SolvedRate        string `json:"solvedRate"`
		DealAgainRate     string `json:"dealAgainRate"`
		ConsumerScore     string `json:"consumerScore"`
		FinalScore        string `json:"finalScore"`
		Status            string `json:"status"`
		StatusDescription string `json:"statusDescription"`
	} `json:"reputation"`
	Phones []struct {
		Number      string        `json:"number"`
		Title       string        `json:"title"`
		WorkingHour []interface{} `json:"workingHour"`
	} `json:"phones"`
	MainTags struct {
		Categories []struct {
			ID          string  `json:"id"`
			Description string  `json:"description"`
			Count       int     `json:"count"`
			Proportion  float64 `json:"proportion"`
		} `json:"categories"`
		Problems []struct {
			ID          string  `json:"id"`
			Description string  `json:"description"`
			Count       int     `json:"count"`
			Proportion  float64 `json:"proportion"`
		} `json:"problems"`
		Products []struct {
			ID          string  `json:"id"`
			Description string  `json:"description"`
			Count       int     `json:"count"`
			Proportion  float64 `json:"proportion"`
		} `json:"products"`
	} `json:"mainTags"`
	PageViews   int `json:"pageViews"`
	CrisisAlert struct {
	} `json:"crisisAlert"`
	ReparoApp struct {
		Enabled bool   `json:"enabled"`
		Link    string `json:"link"`
	} `json:"reparoApp"`
	LeadButton struct {
	} `json:"leadButton"`
	LeadButtonComplain struct {
	} `json:"leadButtonComplain"`
	LeadButtonVideo struct {
	} `json:"leadButtonVideo"`
	SocialNetworks struct {
	} `json:"socialNetworks"`
	Shortname string `json:"shortname"`
	Flags     struct {
		HasNotVerified   bool `json:"hasNotVerified"`
		HasVerified      bool `json:"hasVerified"`
		HasRAFormsLegacy bool `json:"hasRAFormsLegacy"`
		HasRAFormsV2     bool `json:"hasRAFormsV2"`
	} `json:"flags"`
	SinceTime string `json:"sinceTime"`
	Address   struct {
	} `json:"address"`
	CardsCarousel []struct {
		Type        string `json:"type"`
		CardTitle   string `json:"cardTitle"`
		Order       int    `json:"order"`
		PeriodType  string `json:"periodType"`
		Value       string `json:"value"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"cardsCarousel"`
}

type DescribeComplaintResponse struct {
	ID          string      `json:"id"`
	LegacyID    int         `json:"legacyId"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Solved      bool        `json:"solved"`
	Recommend   interface{} `json:"recommend"`
	Status      string      `json:"status"`
	Score       string      `json:"score"`
	Evaluated   bool        `json:"evaluated"`
	Created     string      `json:"created"`
	Category    struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"category"`
	Product struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"product"`
	Problem struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"problem"`
	Company struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Shortname string `json:"shortname"`
		Logo      string `json:"logo"`
	} `json:"company"`
	Interactions []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		Type    string `json:"type"`
		Created string `json:"created"`
		Deleted bool   `json:"deleted"`
	} `json:"interactions"`
	ComplaintsSimiliar []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Created string `json:"created"`
		Company struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Logo      string `json:"logo"`
			Shortname string `json:"shortname"`
		} `json:"company"`
	} `json:"complaintsSimiliar"`
}
