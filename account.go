package socialsso

const (
	// OriginGoogle is the Google origin
	OriginGoogle = Origin(0)
)

type (
	Account struct {

		// ID is the account's unique identifier
		ID			string		`json:"id"`

		// Email is the account's email address
		Email 		string		`json:"email"`

		// Origin is describes the provider that was
		// used when signing in to the application
		Origin Origin `json:"origin"`

		// Token represents a JWT token used for accessing
		// restricted resources etc.
		Token		*JWT `json:"jwt"`
	}

	JWT struct {

		AccessToken 	string		`json:"access_token"`

		RefreshToken 	*string		`json:"refresh_token"`
	}

	Origin uint8
)

// NewGoogleAccount creates and initializes a new Google based account
func NewGoogleAccount(email string, token *JWT) *Account {
	return &Account{Email: email, Origin: OriginGoogle, Token:  nil}
}
