package google

import "github.com/dgrijalva/jwt-go"

type Claims struct {

	Email string `json:"email"`

	Verified bool `json:"verified"`

	FirstName *string `json:"first_name"`

	LastName *string `json:"last_name"`

	jwt.StandardClaims
}
