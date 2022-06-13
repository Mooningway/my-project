package u_rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

func GenerateKey(bits int) (privateKey, publicKey, errMsg string) {
	// generate private key
	key1, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		errMsg = fmt.Sprintf(`RSA failed to generate private key: %v`, err)
		return
	}

	// pkcs1
	var key1Bytes, key2Bytes []byte
	key1Bytes = x509.MarshalPKCS1PrivateKey(key1)

	// public key
	key2 := &key1.PublicKey
	key2Bytes, err = x509.MarshalPKIXPublicKey(key2)
	if err != nil {
		errMsg = fmt.Sprintf(`RSA failed to generate public key: %v`, err)
		return
	}

	// final result
	privateKey = string(pem.EncodeToMemory(&pem.Block{Type: `RSA PRIVATE KEY`, Bytes: key1Bytes}))
	publicKey = string(pem.EncodeToMemory(&pem.Block{Type: `PUBLIC KEY`, Bytes: key2Bytes}))
	return
}

func EncodePkcs1(data, publicKey []byte) (hexString, b64String, errMsg string) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		errMsg = `Rsa public key error.`
		return
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		errMsg = fmt.Sprintf(`Parse public key error: %v`, err)
		return
	}

	pub := pubInterface.(*rsa.PublicKey)

	partLen := pub.N.BitLen()/8 - 11
	chunks := split(data, partLen)

	buffer := bytes.NewBufferString(``)
	for _, c := range chunks {
		cipherResult, err := rsa.EncryptPKCS1v15(rand.Reader, pub, c)
		if err != nil {
			errMsg = fmt.Sprintf(`Encrpt error: %v`, err)
			return
		}

		buffer.Write(cipherResult)
	}

	hexString = hex.EncodeToString(buffer.Bytes())
	b64String = base64.StdEncoding.EncodeToString(buffer.Bytes())
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

func DecodePkcs1B64(cipherText, privateKey []byte) (result, errMsg string) {
	cipherText1, err := base64.StdEncoding.DecodeString(string(cipherText))
	if err != nil {
		errMsg = fmt.Sprintf(`Decrypt error: %v`, err)
		return
	}
	return decodePkcs1(cipherText1, privateKey)
}

func DecodePkcs1Hex(cipherText, privateKey []byte) (result, errMsg string) {
	cipherText1, _ := hex.DecodeString(string(cipherText))
	return decodePkcs1([]byte(cipherText1), privateKey)
}

func decodePkcs1(cipherText, privateKey []byte) (result, errMsg string) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		errMsg = `Rsa private key error.`
		return
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		errMsg = fmt.Sprintf(`Parse private key error: %v`, err)
		return

	}

	partLen := pri.PublicKey.N.BitLen() / 8
	chunks := split(cipherText, partLen)

	buffer := bytes.NewBufferString(``)
	for _, c := range chunks {
		dataResult, err := rsa.DecryptPKCS1v15(rand.Reader, pri, c)
		if err != nil {
			errMsg = fmt.Sprintf(`Decrypt error: %v`, err)
			return
		}

		buffer.Write(dataResult)
	}

	result = string(buffer.String())
	return
}
