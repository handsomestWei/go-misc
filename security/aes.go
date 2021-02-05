package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// aes cbc模式 加密
func AesCbcEncrypt(origData, key, iv []byte) ([]byte, error) {
	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryptData := make([]byte, len(origData))
	// 块加密
	blockMode.CryptBlocks(cryptData, origData)
	return cryptData, nil
}

// aes cbc模式 解密
func AesCbcDecrypt(cryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("AesCbcDecrypt panic %v", err)
			return
		}
	}()

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(origData, cryptData)
	origData = unPadding(origData)
	return origData, nil
}

// 填充
func padding(cipherData []byte, blockSize int) []byte {
	padding := blockSize - len(cipherData)%blockSize
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherData, padData...)
}

func unPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:length-unPadding]
}

// 0填充
func zeroPadding(cipherData []byte, blockSize int) []byte {
	padding := blockSize - len(cipherData)%blockSize
	padData := bytes.Repeat([]byte{0}, padding)
	return append(cipherData, padData...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
