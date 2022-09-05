package u_aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// CBC encrypt
func AesCbcEncrypt(key, originalText, padding, outputEncoding string) (ciperText string, err error) {
	fmt.Println(key)
	fmt.Println(originalText)
	fmt.Println(padding)
	fmt.Println(outputEncoding)

	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		err = newError(ERR_ENCRYPT, err)
		return
	}

	blockSize := block.BlockSize()
	originalBytes := []byte(originalText)

	originalBytes = PKCS7Padding(originalBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])

	encrypts := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(encrypts, originalBytes)

	if OUT_PUT_BSAE64 == strings.ToLower(outputEncoding) {
		// Base64
		ciperText = base64.RawURLEncoding.EncodeToString(encrypts)
	} else {
		// Hex
		ciperText = hex.EncodeToString(encrypts)
	}
	return
}

// CBC decrypt
func AesCbcDecrypt(key, ciperText, padding, outputEncoding string) (originalText string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}

	var ciperTextBytes []byte
	if OUT_PUT_BSAE64 == strings.ToLower(outputEncoding) {
		// Base64
		ciperTextBytes, _ = base64.RawStdEncoding.DecodeString(ciperText)
	} else {
		// Hex
		ciperTextBytes, _ = hex.DecodeString(ciperText)
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	ciperTextBlock := make([]byte, len(ciperTextBytes))
	blockMode.CryptBlocks(ciperTextBlock, ciperTextBytes)

	originalBytes := PKCS7UnPadding(ciperTextBlock)
	originalText = string(originalBytes)
	return
}
