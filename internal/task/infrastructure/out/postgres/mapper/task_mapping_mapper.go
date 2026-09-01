package mapper

import (
	sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/task/domain/model"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/entity"
)

func ToTaskMappingEntity(taskMapping model.TaskMapping) entity.TaskMappingEntity {
	return entity.TaskMappingEntity{
		ID:                   taskMapping.ID,
		OwnerID:              taskMapping.OwnerID,
		ProjectLabel:         taskMapping.ProjectLabel,
		Pattern:              taskMapping.Pattern,
		MatchKeywords:        taskMapping.MatchKeywords,
		MatchOrganizerDomain: taskMapping.MatchOrganizerDomain,
		IssueKey:             taskMapping.IssueKey,
		TypeKey:              taskMapping.TypeKey,
		Notes:                taskMapping.Notes,
		AuditEntity: sharedEntity.AuditEntity{
			CreatedAt: taskMapping.CreatedAt,
			CreatedBy: taskMapping.CreatedBy,
			UpdatedAt: taskMapping.UpdatedAt,
			UpdatedBy: taskMapping.UpdatedBy,
		},
	}
}

func ToTaskMappingModel(taskMappingEntity entity.TaskMappingEntity) model.TaskMapping {
	return model.TaskMapping{
		ID:                   taskMappingEntity.ID,
		OwnerID:              taskMappingEntity.OwnerID,
		ProjectLabel:         taskMappingEntity.ProjectLabel,
		Pattern:              taskMappingEntity.Pattern,
		MatchKeywords:        taskMappingEntity.MatchKeywords,
		MatchOrganizerDomain: taskMappingEntity.MatchOrganizerDomain,
		IssueKey:             taskMappingEntity.IssueKey,
		TypeKey:              taskMappingEntity.TypeKey,
		Notes:                taskMappingEntity.Notes,
		Audit: sharedModel.Audit{
			CreatedAt: taskMappingEntity.CreatedAt,
			CreatedBy: taskMappingEntity.CreatedBy,
			UpdatedAt: taskMappingEntity.UpdatedAt,
			UpdatedBy: taskMappingEntity.UpdatedBy,
		},
	}
}

func ToTaskMappingModelList(taskMappingEntities []entity.TaskMappingEntity) []model.TaskMapping {
	taskMappings := make([]model.TaskMapping, 0, len(taskMappingEntities))
	for _, taskMappingEntity := range taskMappingEntities {
		taskMappings = append(taskMappings, ToTaskMappingModel(taskMappingEntity))
	}

	return taskMappings
}
