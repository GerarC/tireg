package model

import sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

type GlossaryProject struct {
	ID             string
	OwnerID        string
	ProjectLabel   string
	Client         string
	JiraProjectKey string
	BoardURL       string
	Notes          string
	sharedModel.Audit
}
