package u_hex

import (
	"encoding/hex"
	"log"
)

func Encode(val string) string {
	return hex.EncodeToString([]byte(val))
}

func Decode(val string) string {
	result, err := hex.DecodeString(val)
	if err != nil {
		log.Panicln(err)
		return ""
	}
	return string(result)
}
