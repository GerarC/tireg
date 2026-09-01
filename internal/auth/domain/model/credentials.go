package model

import "time"

type Credentials struct {
	Identifier string
	Password   string
}

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
}
