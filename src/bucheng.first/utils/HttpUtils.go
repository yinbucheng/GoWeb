package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get Error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return ""
	}
	return string(body)
}

func PostJson(url string, req string) string {
	temp := bytes.NewBuffer([]byte(req))
	body_type := "application/json;charset=utf-8"
	resp, err := http.Post(url, body_type, temp)
	if err != nil {
		fmt.Println("Post Error:", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return ""
	}
	return string(body)
}

func GetWithHead(url string, heads map[string]string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("New Request Error:", err)
		return ""
	}
	if heads != nil {
		for e := range heads {
			request.Header.Set(e, heads[e])
		}
	}
	request.Header.Set("Connection", "keep-alive")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Get error:", err)
		return ""
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Read Error:", err)
			return ""
		}
		return string(body)
	} else {
		fmt.Println("Response code:", response.StatusCode)
		return ""
	}
}

func PosJsonWithHead(url string, heads map[string]string, req string) string {
	client := &http.Client{}
	temp := bytes.NewBuffer([]byte(req))
	request, err := http.NewRequest("POST", url, temp)
	if err != nil {
		fmt.Println("New Request Error:", err)
		return ""
	}
	request.Header.Set("Content-type", "application/json;charset=utf-8")
	if heads != nil {
		for e := range heads {
			request.Header.Set(e, heads[e])
		}
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Get error:", err)
		return ""
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Read Error:", err)
			return ""
		}
		return string(body)
	} else {
		fmt.Println("Response code:", response.StatusCode)
		return ""
	}
}
