package socialsso

import (
	"context"
	"socialsso/provider/google"
)

type (
	// Service interface that every social SSO provider needs to conform to
	Service interface {
		SignIn(context.Context, string) (*Account, error)
	}

	googleService struct {
		// clientID this is a unique client ID from Google
		clientID	string

		// repository is where the user will be persisted in the end
		// client applications are required to conform to this interface
		repository	Repository
	}
)

// NewGoogleService creates a new Google service instance
func NewGoogleService(clientID string, repository Repository) Service {
	return googleService{clientID: clientID, repository: repository}
}

// SignIn signs the user in to the application using Google
func (s googleService)SignIn(ctx context.Context, idToken string) (*Account, error) {
	claims, err := google.ValidateGoogleJWT(idToken, s.clientID)
	if err != nil {
		return nil, err
	}
	acc := NewGoogleAccount(claims.Email, nil)
	return s.repository.GetOrCreate(ctx, acc)
}
