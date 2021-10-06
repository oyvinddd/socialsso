package google

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.StandardClaims

	// Email the email for the signed-in user
	Email string `json:"email"`
}
