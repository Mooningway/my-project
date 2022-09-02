package u_aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

/*
	AES + GCM
	AES -> Data Encryption Standard
	GCM -> Galois/Counter Mode
*/

// outEncoding string
func AesGcmEncrypt(key, originalText, nonce, encoding string) (ciperText string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return
	}

	nonceBytes := []byte(nonce)
	originalTextBytes := []byte(originalText)
	ciperTextBytes := aesgcm.Seal(nil, nonceBytes, originalTextBytes, nil)
	fmt.Println(encoding)
	if OUT_PUT_BSAE64 == strings.ToLower(encoding) {
		// Base64
		base64.StdEncoding.EncodeToString(ciperTextBytes)
	} else {
		// HEX
		ciperText = hex.EncodeToString(ciperTextBytes)
	}
	return
}

func AesGcmDecrypt(key, ciperText, nonce, encoding string) (originalText string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return
	}

	var ciperTextBytes []byte
	if OUT_PUT_BSAE64 == strings.ToLower(encoding) {
		// Base64
		ciperTextBytes, err = base64.StdEncoding.DecodeString(ciperText)
	} else {
		// HEX
		ciperTextBytes, err = hex.DecodeString(ciperText)
	}
	if err != nil {
		return
	}

	nonceBytes := []byte(nonce)
	originalTextBytes, err := aesgcm.Open(nil, nonceBytes, ciperTextBytes, nil)
	originalText = string(originalTextBytes)
	return
}
