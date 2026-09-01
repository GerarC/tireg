package usecase

import (
	"github.com/gerarc/tireg/internal/user/domain/api"
	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/domain/spi"
)

type FindUserByIdentifierUseCaseImplemented struct {
	userQueryRepository spi.UserQueryRepository
}

func NewFindUserByIdentifierUseCase(userQueryRepository spi.UserQueryRepository) api.FindUserByIdentifierUseCase {
	return &FindUserByIdentifierUseCaseImplemented{userQueryRepository: userQueryRepository}
}

func (findUserByIdentifierUseCase *FindUserByIdentifierUseCaseImplemented) FindByIdentifier(identifier string) (model.User, error) {
	return findUserByIdentifierUseCase.userQueryRepository.FindByIdentifier(identifier)
}
