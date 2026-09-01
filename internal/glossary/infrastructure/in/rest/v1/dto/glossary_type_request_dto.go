package dto

type GlossaryTypeRequestDTO struct {
	TypeKey     string `json:"type_key"`
	Label       string `json:"label"`
	Description string `json:"description"`
}
