package account

const (
	OriginGoogle = Origin(0)
)

type (
	Account struct {

		ID			string		`json:"id"`

		Email 		string		`json:"email"`

		Token		*JWT		`json:"jwt"`
	}

	JWT struct {

		AccessToken 	string 	`json:"access_token"`

		RefreshToken 	string 	`json:"refresh_token"`

		Origin 			Origin	`json:"origin"`
	}

	Origin uint8
)

func New(id string, email string, jwt *JWT) *Account {
	return &Account{ID: id, Email: email, Token: jwt}
}
