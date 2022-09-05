package u_rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
)

const (
	PKCS1                   string = `PKCS1`
	PKCS8                   string = `PKCS8`
	OUT_PUT_BSAE64          string = `base64`
	OUT_PUT_HEX             string = `hex`
	ERR_PRIVATE_KEY         string = `RSA - private key error: %v`
	ERR_PRIVATE_KEY_SUPPORT string = `RSA - private key error: only support PKCS1 or PKCS8`
	ERR_PUBLIC_KEY          string = `RSA - public key error: %v`
	ERR_ENCRYPT             string = `RSA - encrpt error: %v`
	ERR_DECRYPT             string = `RSA - decrypt error: %v`
)

func GenerateKey(bits int) (rsaPrivateKey *rsa.PrivateKey, rsaPublicKey *rsa.PublicKey, err error) {
	// generate private key and public key
	rsaPrivateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		err = newError(ERR_PRIVATE_KEY, err)
		return
	}
	rsaPublicKey = &rsaPrivateKey.PublicKey
	return
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
