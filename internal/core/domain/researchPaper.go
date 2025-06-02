package domain


type ResearchPaper struct {
	ID string `json:"id"`
	Title string `json:"title"`
	ShortDescription string `json:"short_description"`
	Link string `json:"link"`
	ResearchTeam string `json:"research_team"`
	PublishedAt string `json:"published_at"`
	ResearchCategory 
}