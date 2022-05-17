package u_file

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJson(filePath string, data interface{}) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return json.Unmarshal(fileBytes, data)
}

func OverWriteFormatJson(filePath string, data interface{}) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var buffer bytes.Buffer
	jsonBytes, _ := json.Marshal(data)
	json.Indent(&buffer, jsonBytes, ``, `    `)
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
