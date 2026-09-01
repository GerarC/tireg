package model

import sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

type TaskMapping struct {
	ID                   string
	OwnerID              string
	ProjectLabel         string
	Pattern              string
	MatchKeywords        []string
	MatchOrganizerDomain string
	IssueKey             string
	TypeKey              string
	Notes                string
	sharedModel.Audit
}
