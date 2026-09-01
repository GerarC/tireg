package usecase

import (
	"strings"

	sharedSpi "github.com/gerarc/tireg/internal/shared/domain/spi"

	userApi "github.com/gerarc/tireg/internal/user/domain/api"

	"github.com/gerarc/tireg/internal/auth/domain/api"
	"github.com/gerarc/tireg/internal/auth/domain/exception"
	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/domain/spi"
	"github.com/gerarc/tireg/internal/auth/domain/util/constant"
)

type LoginUseCaseImplemented struct {
	findUserByIdentifierUseCase userApi.FindUserByIdentifierUseCase
	passwordHasher              sharedSpi.PasswordHasher
	tokenIssuer                 spi.TokenIssuer
}

func NewLoginUseCase(findUserByIdentifierUseCase userApi.FindUserByIdentifierUseCase, passwordHasher sharedSpi.PasswordHasher, tokenIssuer spi.TokenIssuer) api.LoginUseCase {
	return &LoginUseCaseImplemented{
		findUserByIdentifierUseCase: findUserByIdentifierUseCase,
		passwordHasher:              passwordHasher,
		tokenIssuer:                 tokenIssuer,
	}
}

func (loginUseCase *LoginUseCaseImplemented) Login(credentials model.Credentials) (model.AccessToken, error) {
	var details []string

	if strings.TrimSpace(credentials.Identifier) == "" {
		details = append(details, constant.DetailIdentifierRequired)
	}

	if strings.TrimSpace(credentials.Password) == "" {
		details = append(details, constant.DetailPasswordRequired)
	}

	if len(details) > 0 {
		return model.AccessToken{}, exception.NewValidationFailedError(details...)
	}

	user, err := loginUseCase.findUserByIdentifierUseCase.FindByIdentifier(credentials.Identifier)
	if err != nil {
		return model.AccessToken{}, exception.NewInvalidCredentialsError()
	}

	if !loginUseCase.passwordHasher.Verify(user.PasswordHash, credentials.Password) {
		return model.AccessToken{}, exception.NewInvalidCredentialsError()
	}

	return loginUseCase.tokenIssuer.Issue(user.ID, user.Username)
}
