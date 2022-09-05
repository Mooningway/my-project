package u_aes

/*
	AES (Data Encryption Standard)

	key length:
	16 = aes-128
	24 = aes-192
	32 = aes-256
*/

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	OUT_PUT_BSAE64 string = `base64`
	OUT_PUT_HEX    string = `hex`
	ERR_ENCRYPT    string = `AES - encrpt error: %v`
	ERR_DECRYPT    string = `AES - decrypt error: %v`
)

// AES encryption padding
// type Padding int

// const (
// 	ZERO Padding = 1 + iota
// 	PKCS5
// 	PKCS7
// )

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

func newError(msg string, args ...interface{}) error {
	var errmsg string
	if len(args) == 0 {
		errmsg = msg
	} else {
		errmsg = fmt.Sprintf(msg, args...)
	}
	return errors.New(errmsg)
}
