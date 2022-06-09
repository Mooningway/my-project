package u_md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Hex(value []byte) string {
	hash := md5.New()
	hash.Write(value)
	return hex.EncodeToString(hash.Sum(nil))
}

func HexString(value string) string {
	return Hex([]byte(value))
}
