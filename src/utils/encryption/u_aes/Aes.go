package u_aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

const (
	KEY_LENGTH_16_24_32 string = `the key length must be 16/24/32`
)

// AES encryption padding
type Padding int

const (
	ZERO Padding = 1 + iota
	PKCS5
	PKCS7
)

// AES Output
type Output int

const (
	Hex Output = 1 + iota
	Base64
)

// Padding PKCS7
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

// UnPadding PKCS7
func PKCS7UnPadding(originalText []byte) []byte {
	length := len(originalText)
	unpadding := int(originalText[length-1])
	return originalText[:(length - unpadding)]
}

// CBC encrypt
func CbcEncrypt(original, key string, padding, output int) (result string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		err = errors.New(KEY_LENGTH_16_24_32)
		return
	}

	blockSize := block.BlockSize()
	originalBytes := []byte(original)

	switch padding {
	default:
		originalBytes = PKCS7Padding(originalBytes, blockSize)
	}
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])

	encrypts := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(encrypts, originalBytes)

	switch output {
	case int(Base64):
		result = base64.RawURLEncoding.EncodeToString(encrypts)
	default:
		result = hex.EncodeToString(encrypts)
	}
	return
}

// CBC decrypt
func CbcDecrypt(encrypts, key string, padding, output int) (result string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		err = errors.New(KEY_LENGTH_16_24_32)
		return
	}

	var encryptsBytes []byte
	switch output {
	case int(Base64):
		encryptsBytes, _ = base64.RawStdEncoding.DecodeString(encrypts)
	default:
		encryptsBytes, _ = hex.DecodeString(encrypts)
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	decrypts := make([]byte, len(encryptsBytes))
	blockMode.CryptBlocks(decrypts, encryptsBytes)

	switch padding {
	default:
		decrypts = PKCS7UnPadding(decrypts)
		result = string(decrypts)
	}
	return
}
