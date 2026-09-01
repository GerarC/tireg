package usecase

import (
	"net/mail"
	"strings"

	sharedSpi "github.com/gerarc/tireg/internal/shared/domain/spi"

	"github.com/gerarc/tireg/internal/user/domain/api"
	"github.com/gerarc/tireg/internal/user/domain/exception"
	"github.com/gerarc/tireg/internal/user/domain/model"
	"github.com/gerarc/tireg/internal/user/domain/spi"
	"github.com/gerarc/tireg/internal/user/domain/util/constant"
)

type RegisterUserUseCaseImplemented struct {
	userCommandRepository spi.UserCommandRepository
	userQueryRepository   spi.UserQueryRepository
	passwordHasher        sharedSpi.PasswordHasher
}

func NewRegisterUserUseCase(userCommandRepository spi.UserCommandRepository, userQueryRepository spi.UserQueryRepository, passwordHasher sharedSpi.PasswordHasher) api.RegisterUserUseCase {
	return &RegisterUserUseCaseImplemented{
		userCommandRepository: userCommandRepository,
		userQueryRepository:   userQueryRepository,
		passwordHasher:        passwordHasher,
	}
}

func (registerUserUseCase *RegisterUserUseCaseImplemented) Register(registration model.UserRegistration) (model.User, error) {
	if details := validateRegistration(registration); len(details) > 0 {
		return model.User{}, exception.NewValidationFailedError(details...)
	}

	var conflictDetails []string

	usernameExists, err := registerUserUseCase.userQueryRepository.ExistsByUsername(registration.Username)
	if err != nil {
		return model.User{}, err
	}
	if usernameExists {
		conflictDetails = append(conflictDetails, constant.DetailUsernameAlreadyTaken)
	}

	emailExists, err := registerUserUseCase.userQueryRepository.ExistsByEmail(registration.Email)
	if err != nil {
		return model.User{}, err
	}
	if emailExists {
		conflictDetails = append(conflictDetails, constant.DetailEmailAlreadyTaken)
	}

	if len(conflictDetails) > 0 {
		return model.User{}, exception.NewAlreadyTakenError(conflictDetails...)
	}

	passwordHash, err := registerUserUseCase.passwordHasher.Hash(registration.Password)
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		FirstName:    registration.FirstName,
		LastName:     registration.LastName,
		Username:     registration.Username,
		Email:        registration.Email,
		PasswordHash: passwordHash,
	}
	newUser.CreatedBy = constant.SelfRegistrationActor
	newUser.UpdatedBy = constant.SelfRegistrationActor

	return registerUserUseCase.userCommandRepository.Save(newUser)
}

func validateRegistration(registration model.UserRegistration) []string {
	var details []string

	if strings.TrimSpace(registration.FirstName) == "" {
		details = append(details, constant.DetailFirstNameRequired)
	}

	if strings.TrimSpace(registration.LastName) == "" {
		details = append(details, constant.DetailLastNameRequired)
	}

	if len(registration.Username) < constant.MinUsernameLength {
		details = append(details, constant.DetailUsernameTooShort)
	}

	if _, err := mail.ParseAddress(registration.Email); err != nil {
		details = append(details, constant.DetailInvalidEmail)
	}

	if len(registration.Password) < constant.MinPasswordLength {
		details = append(details, constant.DetailPasswordTooShort)
	}

	return details
}
