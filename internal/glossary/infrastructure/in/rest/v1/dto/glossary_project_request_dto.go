package dto

type GlossaryProjectRequestDTO struct {
	ProjectLabel   string `json:"project_label"`
	Client         string `json:"client"`
	JiraProjectKey string `json:"jira_project_key"`
	BoardURL       string `json:"board_url"`
	Notes          string `json:"notes"`
}
