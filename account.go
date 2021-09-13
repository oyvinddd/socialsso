package socialsso

const (
	// OriginGoogle means the user signed in using Google
	OriginGoogle = Origin(0)

	// FacebookOrigin means the user signed in using Facebook
	FacebookOrigin = Origin(1)

	// AppleOrigin means the user signed in using Apple
	AppleOrigin = Origin(2)
)

type (
	// Account represents a user account in the client application
	Account struct {

		// ID is the account's unique identifier
		ID string `json:"id"`

		// Email is the account's email address
		Email string `json:"email"`

		// Origin is describes the provider that was
		// used when signing in to the application
		Origin Origin `json:"origin"`

		// Token represents a JWT token used for accessing
		// restricted resources in the application
		Token *JWT `json:"jwt"`
	}

	// JWT wraps tokens
	JWT struct {

		// AccessToken used for getting access to secure resources
		AccessToken string `json:"access_token"`

		// RefreshToken used for refreshing an expired access token
		RefreshToken *string `json:"refresh_token"`
	}

	// Origin specifies the provider that was used when authenticating
	Origin uint8
)

// NewGoogleAccount creates and initializes a new Google based account
func NewGoogleAccount(email string, token *JWT) *Account {
	return &Account{Email: email, Origin: OriginGoogle, Token:  token}
}
