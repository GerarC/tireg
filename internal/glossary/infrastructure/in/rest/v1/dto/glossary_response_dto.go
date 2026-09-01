package dto

type GlossaryResponseDTO struct {
	Types    []GlossaryTypeResponseDTO    `json:"types"`
	Projects []GlossaryProjectResponseDTO `json:"projects"`
}
