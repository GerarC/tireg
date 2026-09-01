package mapper

import (
	"time"

	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/dto"
)

func ToTimeEntry(requestDTO dto.TimeEntryRequestDTO) model.TimeEntry {
	return model.TimeEntry{
		Date:          requestDTO.Date,
		ProjectLabel:  requestDTO.ProjectLabel,
		TypeKey:       requestDTO.TypeKey,
		IssueKey:      requestDTO.IssueKey,
		Start:         requestDTO.Start,
		End:           requestDTO.End,
		Hours:         requestDTO.Hours,
		Description:   requestDTO.Description,
		JiraWorklogID: requestDTO.JiraWorklogID,
	}
}

func ToTimeEntryResponseDTO(timeEntry model.TimeEntry) dto.TimeEntryResponseDTO {
	return dto.TimeEntryResponseDTO{
		ID:            timeEntry.ID,
		Date:          timeEntry.Date,
		ProjectLabel:  timeEntry.ProjectLabel,
		TypeKey:       timeEntry.TypeKey,
		IssueKey:      timeEntry.IssueKey,
		Start:         timeEntry.Start,
		End:           timeEntry.End,
		Hours:         timeEntry.Hours,
		Description:   timeEntry.Description,
		JiraWorklogID: timeEntry.JiraWorklogID,
		CreatedAt:     timeEntry.CreatedAt.Format(time.RFC3339),
		CreatedBy:     timeEntry.CreatedBy,
		UpdatedAt:     timeEntry.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:     timeEntry.UpdatedBy,
	}
}

func ToTimeEntryResponseDTOList(timeEntries []model.TimeEntry) []dto.TimeEntryResponseDTO {
	responseDTOs := make([]dto.TimeEntryResponseDTO, 0, len(timeEntries))
	for _, timeEntry := range timeEntries {
		responseDTOs = append(responseDTOs, ToTimeEntryResponseDTO(timeEntry))
	}

	return responseDTOs
}
