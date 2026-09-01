package dto

type TimeEntryRequestDTO struct {
	Date          string  `json:"date"`
	ProjectLabel  string  `json:"project_label"`
	TypeKey       string  `json:"type_key"`
	IssueKey      string  `json:"issue_key"`
	Start         string  `json:"start"`
	End           string  `json:"end"`
	Hours         float64 `json:"hours"`
	Description   string  `json:"description"`
	JiraWorklogID string  `json:"jira_worklog_id"`
}
