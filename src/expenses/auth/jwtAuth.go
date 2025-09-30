package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(sub int32) (string, error) {
	// TODO - check if we can or should use JTI
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "expenseswebapp",
		"sub": sub,
		"aud": sub,
		// "nbf": time.Now().UTC(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("any?"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// ValidateToken takes a JWT string and validates it. Returns the claims
// if the token is valid (the auth middleware handles the context).
func ValidateToken(tStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tStr, func(_ *jwt.Token) (any, error) {
		return []byte("any?"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, nil
}
