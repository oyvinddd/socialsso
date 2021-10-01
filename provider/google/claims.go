package google

import "github.com/golang-jwt/jwt/v4"

type Claims struct {

	// Email the email for the signed-in user
	Email string `json:"email"`

	jwt.StandardClaims
}
