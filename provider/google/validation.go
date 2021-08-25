package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// URL to public key from Google
	publicKeyURL string = "https://www.googleapis.com/oauth2/v1/certs"

	// Issuer of the ID token
	claimsIssuer string = "accounts.google.com"

	// Issuer of the ID token (2)
	claimsIssuer2 string = "https://accounts.google.com"
)

var (
	errInvalidJWT = errors.New("invalid Google JWT")

	errInvalidAUD = errors.New("AUD is invalid")

	errInvalidISS = errors.New("invalid ISS")

	errExpiredJWT =  errors.New("JWT expired")

	errMissingKey = errors.New("public key not found")
)

// https://qvault.io/2020/07/22/how-to-implement-sign-in-with-google-in-golang/
// TODO: check that my validation follows the steps in the docs:
// https://developers.google.com/identity/sign-in/ios/backend-auth
// ValidateGoogleJWT validates an incoming JWT and checks that it is in fact coming from Google
func ValidateGoogleJWT(idToken, clientID string) (*Claims, error) {
	claims := &Claims{}
	// in this case, we only care about the access token, since that's the one from Google.
	// refresh token field is empty since only access token was sent to us by the user trying
	// to log in using Google.
	token, err := jwt.ParseWithClaims(idToken, claims, func(token *jwt.Token) (interface{}, error) {
		keyID := fmt.Sprintf("%s", token.Header["kid"])
		pem, err := getGooglePublicKey(keyID)
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM([]byte(*pem))
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errInvalidJWT
	}
	if claims.Issuer != claimsIssuer && claims.Issuer != claimsIssuer2 {
		return nil, errInvalidISS
	}
	if claims.Audience != clientID {
		return nil, errInvalidAUD
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errExpiredJWT
	}
	return claims, nil
}

// requests public key from Google
func getGooglePublicKey(keyID string) (*string, error) {
	res, err := http.Get(publicKeyURL)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	myResp := map[string]string{}
	err = json.Unmarshal(data, &myResp)
	if err != nil {
		return nil, err
	}
	key, ok := myResp[keyID]
	if !ok {
		return nil, errMissingKey
	}
	return &key, nil
}
