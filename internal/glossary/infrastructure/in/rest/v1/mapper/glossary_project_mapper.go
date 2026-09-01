package mapper

import (
	"time"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
)

func ToGlossaryProject(requestDTO dto.GlossaryProjectRequestDTO) model.GlossaryProject {
	return model.GlossaryProject{
		ProjectLabel:   requestDTO.ProjectLabel,
		Client:         requestDTO.Client,
		JiraProjectKey: requestDTO.JiraProjectKey,
		BoardURL:       requestDTO.BoardURL,
		Notes:          requestDTO.Notes,
	}
}

func ToGlossaryProjectResponseDTO(glossaryProject model.GlossaryProject) dto.GlossaryProjectResponseDTO {
	return dto.GlossaryProjectResponseDTO{
		ID:             glossaryProject.ID,
		ProjectLabel:   glossaryProject.ProjectLabel,
		Client:         glossaryProject.Client,
		JiraProjectKey: glossaryProject.JiraProjectKey,
		BoardURL:       glossaryProject.BoardURL,
		Notes:          glossaryProject.Notes,
		CreatedAt:      glossaryProject.CreatedAt.Format(time.RFC3339),
		CreatedBy:      glossaryProject.CreatedBy,
		UpdatedAt:      glossaryProject.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:      glossaryProject.UpdatedBy,
	}
}

func ToGlossaryProjectResponseDTOList(glossaryProjects []model.GlossaryProject) []dto.GlossaryProjectResponseDTO {
	responseDTOs := make([]dto.GlossaryProjectResponseDTO, 0, len(glossaryProjects))
	for _, glossaryProject := range glossaryProjects {
		responseDTOs = append(responseDTOs, ToGlossaryProjectResponseDTO(glossaryProject))
	}

	return responseDTOs
}
