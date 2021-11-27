package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	resp, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("Content length is: ", resp.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with Golang",
			"price": 0,
			"platform": "pro.learncodeonline,in"
		}
	`)

	resp, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response from post: ", string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "sarvang")
	data.Add("lasttname", "dave")
	data.Add("email", "sarvang@go.dev")

	resp, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Form content: ", string(content))
}
