package u_rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"strings"
)

/*
	PKCS1 (Public-Key Cryptography Standards (PKCS) #1: RSA Cryptography Specifications)
	PKCS8 (Public-Key Cryptography Standards (PKCS) #8: Private-Key Information Syntax Specification)
*/

// Generate Key

func GenerateKeyX509(bits int, pkcs string) (privateKey, publicKey string, err error) {
	rsaPrivateKey, rsaPublicKey, err := GenerateKey(bits)
	if err != nil {
		return
	}

	var privateKeyBytes, publicKeyBytes []byte
	if strings.ToUpper(pkcs) == PKCS1 {
		// PKCS1
		privateKeyBytes = x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	} else if strings.ToUpper(pkcs) == PKCS8 {
		// PKCS8
		privateKeyBytes, err = x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
		if err != nil {
			err = newError(ERR_PRIVATE_KEY, err)
			return
		}
	} else {
		// not support
		err = newError(ERR_PRIVATE_KEY_SUPPORT)
		return
	}

	publicKeyBytes, err = x509.MarshalPKIXPublicKey(rsaPublicKey)
	if err != nil {
		err = newError(ERR_PUBLIC_KEY, err)
		return
	}

	privateKey = string(pem.EncodeToMemory(&pem.Block{Type: `RSA PRIVATE KEY`, Bytes: privateKeyBytes}))
	publicKey = string(pem.EncodeToMemory(&pem.Block{Type: `PUBLIC KEY`, Bytes: publicKeyBytes}))
	return
}

// Encrypt

func EncryptX509(originalText, publicKey, outEncoding string) (ciperText string, err error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		err = newError(ERR_PUBLIC_KEY, `decode error`)
		return
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err = newError(ERR_PUBLIC_KEY, err)
		return
	}
	rsaPublicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		err = newError(ERR_PUBLIC_KEY, err)
		return
	}

	partLen := rsaPublicKey.N.BitLen()/8 - 11
	chunks := split([]byte(originalText), partLen)

	buffer := bytes.NewBufferString(``)
	for _, c := range chunks {
		cipherText, err1 := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, c)
		if err1 != nil {
			err = newError(ERR_ENCRYPT, err1)
			return
		}
		buffer.Write(cipherText)
	}

	if OUT_PUT_BSAE64 == strings.ToLower(outEncoding) {
		// Base64
		ciperText = hex.EncodeToString(buffer.Bytes())
	} else {
		// Hex
		ciperText = base64.StdEncoding.EncodeToString(buffer.Bytes())
	}
	return
}

// Decrypt

func DecryptX509(cipherText, privateKey, pkcs, outEncoding string) (originalText string, err error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = newError(ERR_PRIVATE_KEY, `decode error`)
		return
	}

	var rsaPrivateKey *rsa.PrivateKey
	if strings.ToUpper(pkcs) == PKCS1 {
		// PKCS1
		rsaPrivateKey1, err1 := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err1 != nil {
			err = newError(ERR_PRIVATE_KEY, err1)
			return
		}
		rsaPrivateKey = rsaPrivateKey1
	} else if strings.ToUpper(pkcs) == PKCS8 {
		// PKCS8
		rsaKey, err1 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err1 != nil {
			err = newError(ERR_PRIVATE_KEY, err1)
			return
		}
		rsaPrivateKey1, ok := rsaKey.(*rsa.PrivateKey)
		if !ok {
			err = newError(ERR_PRIVATE_KEY, err)
			return
		}
		rsaPrivateKey = rsaPrivateKey1
	} else {
		// not support
		err = newError(ERR_PRIVATE_KEY_SUPPORT)
		return
	}

	var cipherTextBytes []byte
	if OUT_PUT_BSAE64 == strings.ToLower(outEncoding) {
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

	partLen := rsaPrivateKey.PublicKey.N.BitLen() / 8
	chunks := split(cipherTextBytes, partLen)

	buffer := bytes.NewBufferString(``)
	for _, c := range chunks {
		dataResult, err1 := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, c)
		if err1 != nil {
			err = newError(ERR_DECRYPT, err1)
			return
		}
		buffer.Write(dataResult)
	}

	originalText = buffer.String()
	return
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	lenBuf := len(buf)
	if lenBuf > 0 {
		chunks = append(chunks, buf[:lenBuf])
	}
	return chunks
}
