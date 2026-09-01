package dto

type HealthResponseDTO struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}
