package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// generates the sha1 hash of a given string
func GenerateSha1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// generates the sha256 hash of a given string
func GenerateSha256(text string) string {
	hasher := sha256.New()

	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

// generates the sha256 hash of a given string
func GenerateSha512(text string) string {
	hasher := sha512.New()

	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

// generates the md5 hash of a given string
func GenerateMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// generates the bcrypt hash of a given string
func GenerateBcrypt(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 1)
	return string(bytes)
}
