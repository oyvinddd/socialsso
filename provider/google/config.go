package google

type Config struct {

	ClientID string

	ClientSecret string

	RedirectURL string

	Scopes []string
}

func NewConfig(clientID, clientSecret, redirectURL string, scopes []string) Config {
	return Config{ClientID: clientID, ClientSecret: clientSecret, RedirectURL: redirectURL, Scopes: scopes}
}
