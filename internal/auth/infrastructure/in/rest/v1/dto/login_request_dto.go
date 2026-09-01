package dto

type LoginRequestDTO struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
