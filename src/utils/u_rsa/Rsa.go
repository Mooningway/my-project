package u_rsa

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"strings"
)

const (
	PKCS1                            string = `PKCS1`
	PKCS8                            string = `PKCS8`
	OUT_PUT_BSAE64                   string = `BASE64`
	OUT_PUT_HEX                      string = `HEX`
	BITS_MAX                         int    = 8192
	BITS_MIN                         int    = 12
	HASH_SHA1                        string = `SHA1`
	HASH_SHA256                      string = `SHA256`
	HASH_SHA512                      string = `SHA512`
	HASH_MD5                         string = `MD5`
	ERR_GENERATE_PRIVATE_KEY         string = `RSA - generate private key error: %v`
	ERR_GENERATE_PRIVATE_KEY_SUPPORT string = `RSA - generate private key error: only support PKCS1 or PKCS8`
	ERR_GENERATE_PUBLIC_KEY          string = `RSA - generate public key error: %v`
	ERR_DECODE_PRIVATE_KEY           string = `RSA - decode private key error`
	ERR_DECODE_PUBLIC_KEY            string = `RSA - decode public key error`
	ERR_ENCRYPT                      string = `RSA - encrpt error: %v`
	ERR_DECRYPT                      string = `RSA - decrypt error: %v`
	ERR_HASH_SUPPORT                 string = `RSA - hash not support`
)

/*
	Generate Key

	Params:
		Bits
			key size, must be greater than 12
		Pkcs
			PKCS1, PKCS8
*/
func GenerateKey(bits int, pkcs string) (privateKey, publicKey string, err error) {
	if bits < BITS_MIN {
		bits = BITS_MIN
	} else if bits > BITS_MAX {
		bits = BITS_MAX
	}
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		err = newError(ERR_GENERATE_PRIVATE_KEY, err)
		return
	}
	rsaPublicKey := &rsaPrivateKey.PublicKey

	var privateKeyBytes, publicKeyBytes []byte

	// Private key
	if strings.ToUpper(pkcs) == PKCS1 {
		// PKCS1
		privateKeyBytes = x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	} else if strings.ToUpper(pkcs) == PKCS8 {
		// PKCS8
		privateKeyBytes, err = x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
		if err != nil {
			return ``, ``, newError(ERR_GENERATE_PRIVATE_KEY, err)
		}
	} else {
		// Not support
		return ``, ``, newError(ERR_GENERATE_PRIVATE_KEY_SUPPORT)
	}

	// Public key
	publicKeyBytes, err = x509.MarshalPKIXPublicKey(rsaPublicKey)
	if err != nil {
		return ``, ``, newError(ERR_GENERATE_PUBLIC_KEY, err)
	}

	return string(pem.EncodeToMemory(&pem.Block{Type: `RSA PRIVATE KEY`, Bytes: privateKeyBytes})), string(pem.EncodeToMemory(&pem.Block{Type: `PUBLIC KEY`, Bytes: publicKeyBytes})), nil
}

func EncryptOaep(originalText, publicKey, textOutEncoding string, hash hash.Hash, label []byte) (ciperText string, err error) {
	rsaPublicKey, err := restoreRsaPublicKey(publicKey)
	if err != nil {
		return
	}
	ciperTextBytes, err := rsa.EncryptOAEP(hash, rand.Reader, rsaPublicKey, []byte(originalText), label)
	if err != nil {
		err = newError(ERR_ENCRYPT, err)
		return
	}
	if strings.ToUpper(textOutEncoding) == OUT_PUT_BSAE64 {
		// Base64
		ciperText = base64.StdEncoding.EncodeToString(ciperTextBytes)
	} else {
		// Hex
		ciperText = hex.EncodeToString(ciperTextBytes)
	}
	return
}

func DecryptOaep(cipherText string, privateKey, textOutEncoding, pkcs string, hash hash.Hash, label []byte) (originalText string, err error) {
	rsaPrivateKey, err := restoreRsaPirvateKey(privateKey, pkcs)
	if err != nil {
		return
	}
	var cipherTextBytes []byte
	if strings.ToLower(textOutEncoding) == OUT_PUT_BSAE64 {
		// Base64
		cipherTextBytes, err = base64.StdEncoding.DecodeString(cipherText)
	} else {
		// HEX
		cipherTextBytes, err = hex.DecodeString(cipherText)
	}
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}
	originalTextBytes, err := rsa.DecryptOAEP(hash, rand.Reader, rsaPrivateKey, cipherTextBytes, label)
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}
	originalText = string(originalTextBytes)
	return
}

func EncryptPKCS1v15(originalText, publicKey, textOutEncoding string) (ciperText string, err error) {
	rsaPublicKey, err := restoreRsaPublicKey(publicKey)
	if err != nil {
		return
	}
	ciperTextBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(originalText))
	if err != nil {
		err = newError(ERR_ENCRYPT, err)
		return
	}
	if strings.ToUpper(textOutEncoding) == OUT_PUT_BSAE64 {
		// Base64
		ciperText = base64.StdEncoding.EncodeToString(ciperTextBytes)
	} else {
		// Hex
		ciperText = hex.EncodeToString(ciperTextBytes)
	}
	return
}

func DecryptPKCS1v15(cipherText string, privateKey, textOutEncoding, pkcs string) (originalText string, err error) {
	rsaPrivateKey, err := restoreRsaPirvateKey(privateKey, pkcs)
	if err != nil {
		return
	}
	var cipherTextBytes []byte
	if strings.ToUpper(textOutEncoding) == OUT_PUT_BSAE64 {
		// Base64
		cipherTextBytes, err = base64.StdEncoding.DecodeString(cipherText)
	} else {
		// HEX
		cipherTextBytes, err = hex.DecodeString(cipherText)
	}
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}
	originalTextBytes, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, cipherTextBytes)
	if err != nil {
		err = newError(ERR_DECRYPT, err)
		return
	}
	originalText = string(originalTextBytes)
	return
}

func NewHash(hashMode string) (hash.Hash, error) {
	switch strings.ToUpper(hashMode) {
	case HASH_SHA1:
		return sha1.New(), nil
	case HASH_SHA256:
		return sha256.New(), nil
	case HASH_SHA512:
		return sha256.New(), nil
	case HASH_MD5:
		return md5.New(), nil
	default:
		return nil, newError(ERR_HASH_SUPPORT)
	}
}

func restoreRsaPublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, newError(ERR_DECODE_PUBLIC_KEY)
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, newError(ERR_DECODE_PUBLIC_KEY+`: %v`, err)
	}
	rsaPublicKey, _ := pub.(*rsa.PublicKey)
	return rsaPublicKey, nil
}

func restoreRsaPirvateKey(privateKey, pkcs string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, newError(ERR_DECODE_PRIVATE_KEY)
	}
	privateKeyBytes := block.Bytes

	// PKCS
	if strings.ToUpper(pkcs) == PKCS1 {
		// PKCS1
		rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
		if err != nil {
			return nil, newError(ERR_DECODE_PRIVATE_KEY+`: %v`, err)
		}
		return rsaPrivateKey, nil
	} else if strings.ToUpper(pkcs) == PKCS8 {
		// PKCS8
		rsaKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
		if err != nil {
			return nil, newError(ERR_DECODE_PRIVATE_KEY+`: %v`, err)
		}
		rsaPrivateKey, _ := rsaKey.(*rsa.PrivateKey)
		return rsaPrivateKey, nil
	} else {
		// Not support
		return nil, newError(ERR_GENERATE_PRIVATE_KEY_SUPPORT)
	}
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
