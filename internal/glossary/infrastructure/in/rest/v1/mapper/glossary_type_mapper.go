package mapper

import (
	"time"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
)

func ToGlossaryType(requestDTO dto.GlossaryTypeRequestDTO) model.GlossaryType {
	return model.GlossaryType{
		TypeKey:     requestDTO.TypeKey,
		Label:       requestDTO.Label,
		Description: requestDTO.Description,
	}
}

func ToGlossaryTypeResponseDTO(glossaryType model.GlossaryType) dto.GlossaryTypeResponseDTO {
	return dto.GlossaryTypeResponseDTO{
		ID:          glossaryType.ID,
		TypeKey:     glossaryType.TypeKey,
		Label:       glossaryType.Label,
		Description: glossaryType.Description,
		CreatedAt:   glossaryType.CreatedAt.Format(time.RFC3339),
		CreatedBy:   glossaryType.CreatedBy,
		UpdatedAt:   glossaryType.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:   glossaryType.UpdatedBy,
	}
}

func ToGlossaryTypeResponseDTOList(glossaryTypes []model.GlossaryType) []dto.GlossaryTypeResponseDTO {
	responseDTOs := make([]dto.GlossaryTypeResponseDTO, 0, len(glossaryTypes))
	for _, glossaryType := range glossaryTypes {
		responseDTOs = append(responseDTOs, ToGlossaryTypeResponseDTO(glossaryType))
	}

	return responseDTOs
}
