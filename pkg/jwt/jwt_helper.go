package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type DecodedToken struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Iat      int    `json:"iat"`
	Exp      int    `json:"exp"`
	Iss      string `json:"iss"`
	Role     string `json:"role"`
}

func GenerateToken(claims *jwt.Token, secret string) (token string) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ = claims.SignedString(hmacSecret)

	return
}

func VerifyToken(token string, secret string, env string) *DecodedToken {
	token = token[7:]
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
		fmt.Println(err)
		return nil
	}
	return &decodedToken
}

func GetCurrentTime() int {
	return int(jwt.TimeFunc().Unix())
}
