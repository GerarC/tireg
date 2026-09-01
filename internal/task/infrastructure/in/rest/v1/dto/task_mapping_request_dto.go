package dto

type TaskMappingRequestDTO struct {
	ProjectLabel         string   `json:"project_label"`
	Pattern              string   `json:"pattern"`
	MatchKeywords        []string `json:"match_keywords"`
	MatchOrganizerDomain string   `json:"match_organizer_domain"`
	IssueKey             string   `json:"issue_key"`
	TypeKey              string   `json:"type_key"`
	Notes                string   `json:"notes"`
}
