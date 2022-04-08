package jwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
)

type DecodedToken struct {
	Iat      int      `json:"iat"`
	Role     []string `json:"role"`
	UserID   string   `json:"userID"`
	Username string   `json:"username"`
	Iss      string   `json:"iss"`
}

func GenerateToken(claims *jwt.Token, secret string) (token string) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ = claims.SignedString(hmacSecret)

	return
}

func VerifyToken(token string, secret string, env string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	err = json.Unmarshal(jsonString, &decodedToken)
	if err != nil {
		return nil
	}

	return &decodedToken
}
