package u_hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(content, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
