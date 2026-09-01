package constant

const (
	GetGlossaryRoutePath = "GET /api/v1/glossary"

	CreateGlossaryTypeRoutePath = "POST /api/v1/glossary/types"
	ListGlossaryTypesRoutePath  = "GET /api/v1/glossary/types"
	UpdateGlossaryTypeRoutePath = "PUT /api/v1/glossary/types/{id}"
	DeleteGlossaryTypeRoutePath = "DELETE /api/v1/glossary/types/{id}"

	CreateGlossaryProjectRoutePath = "POST /api/v1/glossary/projects"
	ListGlossaryProjectsRoutePath  = "GET /api/v1/glossary/projects"
	UpdateGlossaryProjectRoutePath = "PUT /api/v1/glossary/projects/{id}"
	DeleteGlossaryProjectRoutePath = "DELETE /api/v1/glossary/projects/{id}"
)
