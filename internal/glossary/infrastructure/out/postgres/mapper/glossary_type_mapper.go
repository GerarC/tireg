package mapper

import (
	sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
)

func ToGlossaryTypeEntity(glossaryType model.GlossaryType) entity.GlossaryTypeEntity {
	return entity.GlossaryTypeEntity{
		ID:          glossaryType.ID,
		OwnerID:     glossaryType.OwnerID,
		TypeKey:     glossaryType.TypeKey,
		Label:       glossaryType.Label,
		Description: glossaryType.Description,
		AuditEntity: sharedEntity.AuditEntity{
			CreatedAt: glossaryType.CreatedAt,
			CreatedBy: glossaryType.CreatedBy,
			UpdatedAt: glossaryType.UpdatedAt,
			UpdatedBy: glossaryType.UpdatedBy,
		},
	}
}

func ToGlossaryTypeModel(glossaryTypeEntity entity.GlossaryTypeEntity) model.GlossaryType {
	return model.GlossaryType{
		ID:          glossaryTypeEntity.ID,
		OwnerID:     glossaryTypeEntity.OwnerID,
		TypeKey:     glossaryTypeEntity.TypeKey,
		Label:       glossaryTypeEntity.Label,
		Description: glossaryTypeEntity.Description,
		Audit: sharedModel.Audit{
			CreatedAt: glossaryTypeEntity.CreatedAt,
			CreatedBy: glossaryTypeEntity.CreatedBy,
			UpdatedAt: glossaryTypeEntity.UpdatedAt,
			UpdatedBy: glossaryTypeEntity.UpdatedBy,
		},
	}
}

func ToGlossaryTypeModelList(glossaryTypeEntities []entity.GlossaryTypeEntity) []model.GlossaryType {
	glossaryTypes := make([]model.GlossaryType, 0, len(glossaryTypeEntities))
	for _, glossaryTypeEntity := range glossaryTypeEntities {
		glossaryTypes = append(glossaryTypes, ToGlossaryTypeModel(glossaryTypeEntity))
	}

	return glossaryTypes
}
