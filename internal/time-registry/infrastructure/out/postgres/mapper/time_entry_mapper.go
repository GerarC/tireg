package mapper

import (
	sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/entity"
)

func ToTimeEntryEntity(timeEntry model.TimeEntry) entity.TimeEntryEntity {
	return entity.TimeEntryEntity{
		ID:            timeEntry.ID,
		OwnerID:       timeEntry.OwnerID,
		Date:          timeEntry.Date,
		ProjectLabel:  timeEntry.ProjectLabel,
		TypeKey:       timeEntry.TypeKey,
		IssueKey:      timeEntry.IssueKey,
		Start:         timeEntry.Start,
		End:           timeEntry.End,
		Hours:         timeEntry.Hours,
		Description:   timeEntry.Description,
		JiraWorklogID: timeEntry.JiraWorklogID,
		AuditEntity: sharedEntity.AuditEntity{
			CreatedAt: timeEntry.CreatedAt,
			CreatedBy: timeEntry.CreatedBy,
			UpdatedAt: timeEntry.UpdatedAt,
			UpdatedBy: timeEntry.UpdatedBy,
		},
	}
}

func ToTimeEntryModel(timeEntryEntity entity.TimeEntryEntity) model.TimeEntry {
	return model.TimeEntry{
		ID:            timeEntryEntity.ID,
		OwnerID:       timeEntryEntity.OwnerID,
		Date:          timeEntryEntity.Date,
		ProjectLabel:  timeEntryEntity.ProjectLabel,
		TypeKey:       timeEntryEntity.TypeKey,
		IssueKey:      timeEntryEntity.IssueKey,
		Start:         timeEntryEntity.Start,
		End:           timeEntryEntity.End,
		Hours:         timeEntryEntity.Hours,
		Description:   timeEntryEntity.Description,
		JiraWorklogID: timeEntryEntity.JiraWorklogID,
		Audit: sharedModel.Audit{
			CreatedAt: timeEntryEntity.CreatedAt,
			CreatedBy: timeEntryEntity.CreatedBy,
			UpdatedAt: timeEntryEntity.UpdatedAt,
			UpdatedBy: timeEntryEntity.UpdatedBy,
		},
	}
}

func ToTimeEntryModelList(timeEntryEntities []entity.TimeEntryEntity) []model.TimeEntry {
	timeEntries := make([]model.TimeEntry, 0, len(timeEntryEntities))
	for _, timeEntryEntity := range timeEntryEntities {
		timeEntries = append(timeEntries, ToTimeEntryModel(timeEntryEntity))
	}

	return timeEntries
}
