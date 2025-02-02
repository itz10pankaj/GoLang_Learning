package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response type is %T", response)
	defer response.Body.Close()

	data, err := io.ReadAll(response.Request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
