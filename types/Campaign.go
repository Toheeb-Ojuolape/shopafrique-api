package types

type Campaign struct {
	Title       string
	Description string
	Media       string
	MediaType   string
	Budget      string
	Status      string
	Views       int64
	Clicks      int64
	Impressions int64
	StartDate   string
	EndDate     string
	Objective   string
	Audience    map[string]interface{}
}

type CampaignID struct {
	CampaignID string
}
