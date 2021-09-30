package generator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

/* commands system */
// interface example --->
type Generator interface {
	Hash(text string) string
}

// this will not accept the argument if the func is not implemented
func Run_Hash(gen Generator, text string) string {
	return gen.Hash(text)
}

type HashSha1 struct{}

func (gen HashSha1) Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type HashSha256 struct{}

func (gen HashSha256) Hash(text string) string {
	hasher := sha256.New()

	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

type HashSha512 struct{}

func (gen HashSha512) Hash(text string) string {
	hasher := sha512.New()

	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

type HashMd5 struct{}

func (gen HashMd5) Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type HashBcrypt struct{}

func (gen HashBcrypt) Hash(text string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(text), 1)
	return string(bytes)
}

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

func GenerateHash(str string, hashalgo string) string {
	switch hashalgo {
	case "sha1":
		return GenerateSha1(str)
	case "sha256":
		return GenerateSha256(str)
	case "sha512":
		return GenerateSha512(str)
	case "md5":
		return GenerateMD5(str)
	case "bcrypt":
		return GenerateBcrypt(str)
	}
	return "err. hash algo doesn't exist"
}
