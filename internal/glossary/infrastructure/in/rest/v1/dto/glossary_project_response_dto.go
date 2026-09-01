package dto

type GlossaryProjectResponseDTO struct {
	ID             string `json:"id"`
	ProjectLabel   string `json:"project_label"`
	Client         string `json:"client"`
	JiraProjectKey string `json:"jira_project_key"`
	BoardURL       string `json:"board_url"`
	Notes          string `json:"notes"`
	CreatedAt      string `json:"created_at"`
	CreatedBy      string `json:"created_by"`
	UpdatedAt      string `json:"updated_at"`
	UpdatedBy      string `json:"updated_by"`
}
