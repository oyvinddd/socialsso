package google

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/oyvinddd/socialsso/account"
)

const (
	publicKeyURL string = "https://www.googleapis.com/oauth2/v1/certs"

	claimsIssuer string = "accounts.google.com"

	claimsIssuer2 string = "https://accounts.google.com"
)

var (
	errInvalidJWT = errors.New("invalid Google JWT")

	errInvalidAUD = errors.New("AUD is invalid")

	errInvalidISS = errors.New("invalid ISS")

	errExpiredJWT =  errors.New("JWT expired")

	errMissingKey = errors.New("public key not found")
)

type (

	googleService struct {

		// clientID this is a unique client ID from Google
		clientID	string

		// repository is where the user will be persisted in the end
		// client applications are required to conform to this interface
		repository	account.Repository
	}

	Claims struct {

		// Email the email for the signed-in user
		Email		string	`json:"email"`

		//Verified	bool	`json:"verified"`

		//FirstName	*string `json:"first_name"`

		//LastName	*string `json:"last_name"`

		jwt.StandardClaims
	}
)

// NewGoogleService creates a new Google service instance
func NewGoogleService(clientID string, repository account.Repository) account.Service {
	return &googleService{clientID: clientID, repository: repository}
}

// SignIn signs the user in to the application using Google
func (s googleService)SignIn(ctx context.Context, idToken string) (*account.Account, error) {
	claims, err := validateGoogleJWT(idToken, s.clientID)
	if err != nil {
		return nil, err
	}
	acc := account.NewGoogleAccount(claims.Email, nil)
	return s.repository.GetOrCreate(ctx, acc)
}
