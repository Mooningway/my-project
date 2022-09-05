package u_aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

/*
	AES + GCM
	AES -> Data Encryption Standard
	GCM -> Galois/Counter Mode
*/

func AesGcmEncrypt(key, originalText, nonce, outputEncoding string) (ciperText string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		err = newError(ERR_ENCRYPT, err)
		return
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		err = newError(ERR_ENCRYPT, err)
		return
	}

	nonceBytes := []byte(nonce)
	originalTextBytes := []byte(originalText)
	ciperTextBytes := aesgcm.Seal(nil, nonceBytes, originalTextBytes, nil)
	if OUT_PUT_BSAE64 == strings.ToLower(outputEncoding) {
		// Base64
		ciperText = base64.StdEncoding.EncodeToString(ciperTextBytes)
	} else {
		// HEX
		ciperText = hex.EncodeToString(ciperTextBytes)
	}
	return
}

func AesGcmDecrypt(key, ciperText, nonce, outputEncoding string) (originalText string, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}

	var ciperTextBytes []byte
	if OUT_PUT_BSAE64 == strings.ToLower(outputEncoding) {
		// Base64
		ciperTextBytes, err = base64.StdEncoding.DecodeString(ciperText)
	} else {
		// HEX
		ciperTextBytes, err = hex.DecodeString(ciperText)
	}
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}

	nonceBytes := []byte(nonce)
	originalTextBytes, err := aesgcm.Open(nil, nonceBytes, ciperTextBytes, nil)
	originalText = string(originalTextBytes)
	return
}
