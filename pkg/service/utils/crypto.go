package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"slices"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// GetHash - string encryption
func GetHash(text, code string) string {
	var hasher hash.Hash
	codeMap := map[string]hash.Hash{
		"sha256": sha256.New(),
		"sha512": sha512.New(),
	}
	if !slices.Contains([]string{"sha256", "sha512"}, code) {
		code = "sha256"
	}
	hasher = codeMap[code]
	hasher.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}

func CreatePasswordHash(password string) (hash string, err error) {
	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ComparePasswordAndHash(password string, hash string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
