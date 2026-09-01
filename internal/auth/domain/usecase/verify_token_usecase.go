package usecase

import (
	"github.com/gerarc/tireg/internal/auth/domain/api"
	"github.com/gerarc/tireg/internal/auth/domain/exception"
	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/domain/spi"
)

type VerifyTokenUseCaseImplemented struct {
	tokenVerifier spi.TokenVerifier
}

func NewVerifyTokenUseCase(tokenVerifier spi.TokenVerifier) api.VerifyTokenUseCase {
	return &VerifyTokenUseCaseImplemented{tokenVerifier: tokenVerifier}
}

func (verifyTokenUseCase *VerifyTokenUseCaseImplemented) Verify(token string) (model.AuthenticatedUser, error) {
	authenticatedUser, err := verifyTokenUseCase.tokenVerifier.Verify(token)
	if err != nil {
		return model.AuthenticatedUser{}, exception.NewInvalidCredentialsError()
	}

	return authenticatedUser, nil
}
