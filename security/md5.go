package security

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密转16进制字符串
func Md5ToHexString(s string, args ...string) string {
	ciphered := Md5ToBytes(s, args...)
	return hex.EncodeToString(ciphered)
}

// MD5加密转byte数组
func Md5ToBytes(data string, args ...string) []byte {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	if args != nil && len(args) > 0 {
		for _, salt := range args {
			md5Ctx.Write([]byte(salt))
		}
	}
	return md5Ctx.Sum(nil)
}
