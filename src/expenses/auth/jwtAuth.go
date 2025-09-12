package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO - also allow asymetrical keys to be used
func GenerateToken(sub string) (string, error) {
	// TODO - check if we can or should use JTI
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "expenseswebapp",
		"sub": sub,
		"aud": sub,
		// "nbf": time.Now().UTC(),
		"iat": time.Now().UTC(),
		"exp": time.Now().Add(15 * time.Minute).UTC(),
	})

	tokenStr, err := token.SignedString([]byte("any?"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tStr string) (bool, error) {
	token, err := jwt.Parse(tStr, func(token *jwt.Token) (any, error) {
		return []byte("any?"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))	
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		log.Println(claims)
		return true, nil
	}

	return false, nil
}
