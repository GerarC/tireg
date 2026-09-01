package usecase

import (
	"github.com/gerarc/tireg/internal/glossary/domain/api"
	"github.com/gerarc/tireg/internal/glossary/domain/model"
	"github.com/gerarc/tireg/internal/glossary/domain/spi"
)

type ListGlossaryProjectsUseCaseImplemented struct {
	glossaryProjectQueryRepository spi.GlossaryProjectQueryRepository
}

func NewListGlossaryProjectsUseCase(glossaryProjectQueryRepository spi.GlossaryProjectQueryRepository) api.ListGlossaryProjectsUseCase {
	return &ListGlossaryProjectsUseCaseImplemented{glossaryProjectQueryRepository: glossaryProjectQueryRepository}
}

func (listGlossaryProjectsUseCase *ListGlossaryProjectsUseCaseImplemented) List(ownerID string) ([]model.GlossaryProject, error) {
	return listGlossaryProjectsUseCase.glossaryProjectQueryRepository.SelectAllByOwner(ownerID)
}
