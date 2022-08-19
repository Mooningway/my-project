package u_file

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadLine(filePath string) (fileContent []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		input, errRead := reader.ReadString('\n')
		fileContent = append(fileContent, input)
		if errRead == io.EOF {
			return
		}
	}
}

func ReadCsvLine(filePath string) (result [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		line, err0 := reader.Read()
		if len(line) > 0 {
			result = append(result, line)
		}
		if err0 == io.EOF {
			break
		}
		if err0 != nil {
			err = err0
			return
		}
	}
	return
}

func ReadCsvAll(filePath string) (result [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	return csv.NewReader(file).ReadAll()
}

func ReadAllForJson(filePath string, data interface{}) (err error) {
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

func Write(filePath string, fileContent string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprint(fileContent))
	write.Flush()
	return nil
}

func WriteLine(filePath string, content []string) (err error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for _, val := range content {
		fmt.Fprint(write, val)
	}
	write.Flush()
	return
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
