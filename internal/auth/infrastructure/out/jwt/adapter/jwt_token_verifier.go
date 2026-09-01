package adapter

import (
	"errors"

	jwtlib "github.com/golang-jwt/jwt/v5"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"

	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/domain/spi"
)

type JWTTokenVerifier struct {
	secret []byte
}

func NewJWTTokenVerifier(appConfig *config.Config) spi.TokenVerifier {
	return &JWTTokenVerifier{secret: []byte(appConfig.JWTSecret)}
}

func (jwtTokenVerifier *JWTTokenVerifier) Verify(token string) (model.AuthenticatedUser, error) {
	claims := &jwtClaims{}

	parsedToken, err := jwtlib.ParseWithClaims(token, claims, func(t *jwtlib.Token) (any, error) {
		return jwtTokenVerifier.secret, nil
	}, jwtlib.WithValidMethods([]string{jwtlib.SigningMethodHS256.Name}))
	if err != nil {
		return model.AuthenticatedUser{}, err
	}

	if !parsedToken.Valid {
		return model.AuthenticatedUser{}, errors.New("invalid token")
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return model.AuthenticatedUser{}, err
	}

	return model.AuthenticatedUser{ID: userID, Username: claims.Username}, nil
}
