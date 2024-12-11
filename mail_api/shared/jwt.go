package shared

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	jwt.MapClaims        // expity, IssueAt, etc...
	Session       string `json:"session"`
}

func GetTokenFromRequest(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	tokenSlice := strings.Split(bearerToken, " ")
	if len(tokenSlice) == 2 {
		return tokenSlice[1]
	}
	return ""
}

func ParseJWTToken(tokenStr string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
