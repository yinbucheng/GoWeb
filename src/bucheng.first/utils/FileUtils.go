package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FileReadText(name string) string {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Read File Error")
		return ""
	}
	return string(content)
}

func FileWriteText(name string, content string) bool {
	data := []byte(content)
	if !isFileExist(name) {
		err := ioutil.WriteFile(name, data, 0644)
		if err != nil {
			return false
		}
		return true
	} else {
		f, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("WRITE file fail:", err)
			return false
		}
		defer f.Close()
		f.Write(data)
		return true
	}
}

func WriteMap(name string, data map[string]string) bool {
	content := ""
	for key := range data {
		content += key + "=" + data[key] + ","
	}
	return FileWriteText(name, content)
}

func ReadMap(name string) map[string]string {
	content := FileReadText(name)
	if content == "" {
		return nil
	}
	properties := make(map[string]string)
	temps := strings.FieldsFunc(content, func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	})
	for index := range temps {
		temp := temps[index]
		if temp == "" {
			continue
		}
		temp2 := strings.FieldsFunc(temp, func(r rune) bool {
			if r == '=' {
				return true
			}
			return false
		})
		properties[temp2[0]] = temp2[1]
	}
	return properties
}

func isFileExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}
