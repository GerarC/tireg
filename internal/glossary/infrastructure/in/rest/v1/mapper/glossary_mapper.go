package mapper

import (
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/dto"
)

func ToGlossaryResponseDTO(glossary model.Glossary) dto.GlossaryResponseDTO {
	return dto.GlossaryResponseDTO{
		Types:    ToGlossaryTypeResponseDTOList(glossary.Types),
		Projects: ToGlossaryProjectResponseDTOList(glossary.Projects),
	}
}
