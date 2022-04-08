package hash

import (
	"crypto/sha512"
	"encoding/base64"
)

func GenerateHash(password string, salt string) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}
