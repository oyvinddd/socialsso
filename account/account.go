package account

const (
	// OriginGoogle is the Google origin
	OriginGoogle = Origin(0)
)

type (
	Account struct {

		ID			string		`json:"id"`

		Email 		string		`json:"email"`

		Origin		Origin		`json:"origin"`

		Token		*JWT		`json:"jwt"`
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
