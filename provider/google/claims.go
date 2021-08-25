package google

import "github.com/dgrijalva/jwt-go"

type Claims struct {

	// Email the email for the signed-in user
	Email string `json:"email"`

	jwt.StandardClaims
}
