package structs

type Query struct {
	Calendar   string `json:"calendar"`
	Query      string `json:"query"`
	StartDate  string `json:"dateStart"`
	EndDate    string `json:"dateEnd"`
	MaxResults int64  `json:"maxResults"`
}
