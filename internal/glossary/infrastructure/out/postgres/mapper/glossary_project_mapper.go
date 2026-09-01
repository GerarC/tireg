package mapper

import (
	sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
)

func ToGlossaryProjectEntity(glossaryProject model.GlossaryProject) entity.GlossaryProjectEntity {
	return entity.GlossaryProjectEntity{
		ID:             glossaryProject.ID,
		OwnerID:        glossaryProject.OwnerID,
		ProjectLabel:   glossaryProject.ProjectLabel,
		Client:         glossaryProject.Client,
		JiraProjectKey: glossaryProject.JiraProjectKey,
		BoardURL:       glossaryProject.BoardURL,
		Notes:          glossaryProject.Notes,
		AuditEntity: sharedEntity.AuditEntity{
			CreatedAt: glossaryProject.CreatedAt,
			CreatedBy: glossaryProject.CreatedBy,
			UpdatedAt: glossaryProject.UpdatedAt,
			UpdatedBy: glossaryProject.UpdatedBy,
		},
	}
}

func ToGlossaryProjectModel(glossaryProjectEntity entity.GlossaryProjectEntity) model.GlossaryProject {
	return model.GlossaryProject{
		ID:             glossaryProjectEntity.ID,
		OwnerID:        glossaryProjectEntity.OwnerID,
		ProjectLabel:   glossaryProjectEntity.ProjectLabel,
		Client:         glossaryProjectEntity.Client,
		JiraProjectKey: glossaryProjectEntity.JiraProjectKey,
		BoardURL:       glossaryProjectEntity.BoardURL,
		Notes:          glossaryProjectEntity.Notes,
		Audit: sharedModel.Audit{
			CreatedAt: glossaryProjectEntity.CreatedAt,
			CreatedBy: glossaryProjectEntity.CreatedBy,
			UpdatedAt: glossaryProjectEntity.UpdatedAt,
			UpdatedBy: glossaryProjectEntity.UpdatedBy,
		},
	}
}

func ToGlossaryProjectModelList(glossaryProjectEntities []entity.GlossaryProjectEntity) []model.GlossaryProject {
	glossaryProjects := make([]model.GlossaryProject, 0, len(glossaryProjectEntities))
	for _, glossaryProjectEntity := range glossaryProjectEntities {
		glossaryProjects = append(glossaryProjects, ToGlossaryProjectModel(glossaryProjectEntity))
	}

	return glossaryProjects
}
