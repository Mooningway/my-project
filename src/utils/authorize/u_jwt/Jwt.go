package u_jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type Alg int

const (
	HS256 Alg = 1 + iota
)

var algNames = []string{
	`HS256`,
}

func GetToken(alg Alg, playload map[string]interface{}, secret string) string {
	// Header
	// {alg: "", typ: "JWT"}
	header := make(map[string]string)
	header[`alg`] = alg.String()
	header[`typ`] = `JWT`
	headerJson, _ := json.Marshal(header)
	headerB64 := base64Encode(headerJson)

	// Playload
	playloadJson, _ := json.Marshal(playload)
	playloadB64 := base64Encode(playloadJson)

	// Signature
	var signature string
	switch alg {
	default:
		signature = uhmacSha256(headerB64+`.`+playloadB64, secret)
	}

	// Result
	var result strings.Builder
	result.WriteString(headerB64)
	result.WriteString(`.`)
	result.WriteString(playloadB64)
	result.WriteString(`.`)
	result.WriteString(signature)
	return result.String()
}

func GetTokenHS256(playload map[string]interface{}, secret string) string {
	return GetToken(HS256, playload, secret)
}

func (a Alg) String() string {
	if HS256 <= a && a <= HS256 {
		return algNames[a-1]
	}
	return ``
}

func base64Encode(val []byte) string {
	return base64.StdEncoding.EncodeToString(val)
}

func uhmacSha256(content, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
