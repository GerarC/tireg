package dto

type GlossaryTypeResponseDTO struct {
	ID          string `json:"id"`
	TypeKey     string `json:"type_key"`
	Label       string `json:"label"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
}
