package dto

type TimeEntryResponseDTO struct {
	ID            string  `json:"id"`
	Date          string  `json:"date"`
	ProjectLabel  string  `json:"project_label"`
	TypeKey       string  `json:"type_key"`
	IssueKey      string  `json:"issue_key"`
	Start         string  `json:"start"`
	End           string  `json:"end"`
	Hours         float64 `json:"hours"`
	Description   string  `json:"description"`
	JiraWorklogID string  `json:"jira_worklog_id"`
	CreatedAt     string  `json:"created_at"`
	CreatedBy     string  `json:"created_by"`
	UpdatedAt     string  `json:"updated_at"`
	UpdatedBy     string  `json:"updated_by"`
}
