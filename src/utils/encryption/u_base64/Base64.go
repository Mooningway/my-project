package u_base64

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strings"
)

func Encode(val []byte) string {
	return base64.StdEncoding.EncodeToString(val)
}

func Decode(val []byte) string {
	result, err := base64.StdEncoding.DecodeString(string(val))
	if err != nil {
		log.Println(err)
		return ``
	}
	return string(result)
}

func EncodeFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	fileByteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return ``, err
	}
	var result strings.Builder
	result.WriteString(`data:`)
	result.WriteString(header.Header.Get(`Content-Type`))
	result.WriteString(`;base64,`)
	result.WriteString(Encode(fileByteValue))
	return result.String(), nil
}
