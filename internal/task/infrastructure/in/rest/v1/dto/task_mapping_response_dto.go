package dto

type TaskMappingResponseDTO struct {
	ID                   string   `json:"id"`
	ProjectLabel         string   `json:"project_label"`
	Pattern              string   `json:"pattern"`
	MatchKeywords        []string `json:"match_keywords"`
	MatchOrganizerDomain string   `json:"match_organizer_domain"`
	IssueKey             string   `json:"issue_key"`
	TypeKey              string   `json:"type_key"`
	Notes                string   `json:"notes"`
	CreatedAt            string   `json:"created_at"`
	CreatedBy            string   `json:"created_by"`
	UpdatedAt            string   `json:"updated_at"`
	UpdatedBy            string   `json:"updated_by"`
}
