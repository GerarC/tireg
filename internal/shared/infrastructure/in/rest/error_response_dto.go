package rest

type ErrorResponseDTO struct {
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	Details   []string `json:"details"`
	Timestamp string   `json:"timestamp"`
}
