package google

import (
	"context"
	"errors"
	"soocialsso/account"
)

const (
	publicKeyURL string = "https://www.googleapis.com/oauth2/v1/certs"

	claimsIssuer string = "accounts.google.com"

	claimsIssuer2 string = "https://accounts.google.com"

	clientID string = "81111383572-9s4c8ba9ncg3hjfk6rr2kes5qblrb5qe.apps.googleusercontent.com"
)

var (
	errInvalidJWT = errors.New("invalid Google JWT")

	errInvalidAUD = errors.New("AUD is invalid")

	errInvalidISS = errors.New("invalid ISS")

	errExpiredJWT =  errors.New("JWT expired")

	errMissingKey = errors.New("public key not found")
)

type (
	// TODO: this service should be reused by all social login providers
	Service interface {
		SignIn(ctx context.Context, claims Claims) (*account.Account, error)
	}

	googleService struct {
		repository account.Repository
	}
)

func NewGoogleService(repository account.Repository) Service {
	return &googleService{repository: repository}
}

func (s googleService)SignIn(ctx context.Context, claims Claims) (*account.Account, error) {
	return s.repository.GetOrCreate(ctx, claims.Email)
}
