package crypt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

// https://pkg.go.dev/github.com/dgrijalva/jwt-go#section-readme
var tokenString = "signedJwtTokenString"

func TestParse(t *testing.T) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("hmacSampleSecret"), nil
	})
	t.Log(token)
	t.Log(err)
}
