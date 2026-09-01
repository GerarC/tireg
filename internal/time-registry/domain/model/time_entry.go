package model

import sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

type TimeEntry struct {
	ID            string
	OwnerID       string
	Date          string
	ProjectLabel  string
	TypeKey       string
	IssueKey      string
	Start         string
	End           string
	Hours         float64
	Description   string
	JiraWorklogID string
	sharedModel.Audit
}
