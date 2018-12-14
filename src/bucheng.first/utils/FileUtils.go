package utils

import (
	"fmt"
	"io/ioutil"
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
	err := ioutil.WriteFile(name, data, 0644)
	if err != nil {
		return false
	}
	return true
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
