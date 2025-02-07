package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Requests in go")
	// PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("content length", response.ContentLength)

	// content, _ := io.ReadAll(response.Body)

	// fmt.Println(string(content))

	//second method
	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "https://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"slk":"efmklm"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form Data
	data := url.Values{}
	data.Add("firstName", "Pankaj")

	http.PostForm(myurl, data)

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
