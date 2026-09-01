package mapper

import (
	"time"

	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/dto"
)

func ToTaskMapping(requestDTO dto.TaskMappingRequestDTO) model.TaskMapping {
	return model.TaskMapping{
		ProjectLabel:         requestDTO.ProjectLabel,
		Pattern:              requestDTO.Pattern,
		MatchKeywords:        requestDTO.MatchKeywords,
		MatchOrganizerDomain: requestDTO.MatchOrganizerDomain,
		IssueKey:             requestDTO.IssueKey,
		TypeKey:              requestDTO.TypeKey,
		Notes:                requestDTO.Notes,
	}
}

func ToTaskMappingResponseDTO(taskMapping model.TaskMapping) dto.TaskMappingResponseDTO {
	return dto.TaskMappingResponseDTO{
		ID:                   taskMapping.ID,
		ProjectLabel:         taskMapping.ProjectLabel,
		Pattern:              taskMapping.Pattern,
		MatchKeywords:        taskMapping.MatchKeywords,
		MatchOrganizerDomain: taskMapping.MatchOrganizerDomain,
		IssueKey:             taskMapping.IssueKey,
		TypeKey:              taskMapping.TypeKey,
		Notes:                taskMapping.Notes,
		CreatedAt:            taskMapping.CreatedAt.Format(time.RFC3339),
		CreatedBy:            taskMapping.CreatedBy,
		UpdatedAt:            taskMapping.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:            taskMapping.UpdatedBy,
	}
}

func ToTaskMappingResponseDTOList(taskMappings []model.TaskMapping) []dto.TaskMappingResponseDTO {
	responseDTOs := make([]dto.TaskMappingResponseDTO, 0, len(taskMappings))
	for _, taskMapping := range taskMappings {
		responseDTOs = append(responseDTOs, ToTaskMappingResponseDTO(taskMapping))
	}

	return responseDTOs
}
