package u_file

import (
	"encoding/json"
	"os"
)

// O_WRONLY

func WriteJson(filePath string, jsonBytes []byte) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(string(jsonBytes))
}

func WriteJsonString(filePath string, jsonString string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(jsonString)
}

func ReadJson(filePath string, data interface{}) (err error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}
