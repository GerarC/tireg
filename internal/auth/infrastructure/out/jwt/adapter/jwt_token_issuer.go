package adapter

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"

	"github.com/gerarc/tireg/internal/auth/domain/model"
	"github.com/gerarc/tireg/internal/auth/domain/spi"
)

type jwtClaims struct {
	Username string `json:"username"`
	jwtlib.RegisteredClaims
}

type JWTTokenIssuer struct {
	secret     []byte
	expiration time.Duration
}

func NewJWTTokenIssuer(appConfig *config.Config) spi.TokenIssuer {
	return &JWTTokenIssuer{
		secret:     []byte(appConfig.JWTSecret),
		expiration: appConfig.JWTExpiration,
	}
}

func (jwtTokenIssuer *JWTTokenIssuer) Issue(userID string, username string) (model.AccessToken, error) {
	expiresAt := time.Now().Add(jwtTokenIssuer.expiration)

	claims := jwtClaims{
		Username: username,
		RegisteredClaims: jwtlib.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwtlib.NewNumericDate(expiresAt),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtTokenIssuer.secret)
	if err != nil {
		return model.AccessToken{}, err
	}

	return model.AccessToken{Token: signedToken, ExpiresAt: expiresAt}, nil
}
