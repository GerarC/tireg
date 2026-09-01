package constant

const (
	CodeValidationFailed    = 400
	MessageValidationFailed = "GLOSSARY_VALIDATION_FAILED"

	CodeGlossaryTypeNotFound    = 404
	MessageGlossaryTypeNotFound = "GLOSSARY_TYPE_NOT_FOUND"

	CodeGlossaryProjectNotFound    = 404
	MessageGlossaryProjectNotFound = "GLOSSARY_PROJECT_NOT_FOUND"

	DetailTypeKeyRequired      = "type key is required"
	DetailLabelRequired        = "label is required"
	DetailProjectLabelRequired = "project label is required"
)

type DefaultGlossaryType struct {
	TypeKey     string
	Label       string
	Description string
}

var DefaultGlossaryTypes = []DefaultGlossaryType{
	{TypeKey: "client_meeting", Label: "Reunión con cliente", Description: "Reunión donde el cliente participa u organiza"},
	{TypeKey: "internal_meeting", Label: "Reunión interna", Description: "Reunión solo con el equipo, sin cliente"},
	{TypeKey: "support", Label: "Soporte / bug / gestión puntual", Description: "Atención de incidentes, despliegues o solicitudes puntuales"},
	{TypeKey: "dev", Label: "Desarrollo", Description: "Avance de una historia de usuario o tarea, de cualquier disciplina"},
	{TypeKey: "other", Label: "Otro", Description: "Cualquier cosa que no encaje en las anteriores"},
}
