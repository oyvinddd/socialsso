package socialsso

import (
	"context"
	"github.com/oyvinddd/socialsso/provider/google"
	"golang.org/x/oauth2"
	googl "golang.org/x/oauth2/google"
	"net/url"
	"strings"
)

const (
	googleAuthURL string = "https://accounts.google.com/o/oauth2/auth"

	// NEVER USED
	googleTokenURL string = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

type (
	// Service interface that every social SSO provider needs to conform to
	Service interface {
		// SignIn signs the user into the client application
		SignIn(context.Context, string) (*Account, error)

		// GetRedirectURL creates a redirect URL with the correct config
		// for a given service
		GetRedirectURL() string

		ExchangeCodeAndToken(context.Context, string) (string, error)
	}

	googleService struct {
		// config contains all Google credentials, URLs etc.
		config *oauth2.Config

		// repository is where the user will be persisted in the end
		// client applications are required to conform to this interface
		repository Repository
	}
)

// NewGoogleService creates a new Google service instance
func NewGoogleService(clientID, clientSecret, redirectURL string, scopes []string, repository Repository) Service {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     googl.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}
	return googleService{config: cfg, repository: repository}
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
	URL, _ := url.Parse(googleAuthURL)
	oauthStateString := "" // FIXME: add generated state here for preventing xss attack
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

func (s googleService)ExchangeCodeAndToken(ctx context.Context, code string) (string, error) {
	token, err := s.config.Exchange(ctx, code)
	return token.AccessToken, err
}
