package socialsso

import (
	"context"
	"github.com/oyvinddd/socialsso/provider/google"
	"net/url"
	"strings"
)

const (
	authServiceURL string = "https://accounts.google.com/o/oauth2/auth"
)

type (
	// Service interface that every social SSO provider needs to conform to
	Service interface {
		// SignIn signs the user into the client application
		SignIn(context.Context, string) (*Account, error)

		// GetRedirectURL creates a redirect URL with the correct config
		// for a given service
		GetRedirectURL() string
	}

	googleService struct {
		// config contains all Google credentials, URLs etc.
		config google.Config

		// repository is where the user will be persisted in the end
		// client applications are required to conform to this interface
		repository Repository
	}
)

// NewGoogleService creates a new Google service instance
func NewGoogleService(config google.Config, repository Repository) Service {
	return googleService{config: config, repository: repository}
}

// SignIn signs the user in to the application using Google
func (s googleService)SignIn(ctx context.Context, idToken string) (*Account, error) {
	claims, err := google.ValidateGoogleJWT(idToken, s.config.ClientID)
	if err != nil {
		return nil, err
	}
	acc := NewGoogleAccount(claims.Email, nil)
	return s.repository.GetOrCreate(ctx, acc)
}

func (s googleService)GetRedirectURL() string {
	URL, _ := url.Parse(authServiceURL)
	oauthStateString := "" // FIXME: what is this?
	scopes := s.config.Scopes
	redirectURL := s.config.RedirectURL
	// query parameters
	parameters := url.Values{}
	parameters.Add("client_id", s.config.ClientID)
	parameters.Add("scope", strings.Join(scopes, " "))
	parameters.Add("redirect_uri", redirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	return URL.String()
}
