package security

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// sha1加密
func Sha1(s string) string {
	sha1Ctx := sha1.New()
	sha1Ctx.Write([]byte(s))
	ciphered := sha1Ctx.Sum(nil)
	return hex.EncodeToString(ciphered)
}

// sha256加密
func Sha256(s string) []byte {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(s))
	return sha256Ctx.Sum(nil)
}
