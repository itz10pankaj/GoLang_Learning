package main

import (
	"fmt"
	"net/url"
)

const myurl string = "lkmf"

func main() {
	fmt.Println("URL in golang")
	// fmt.Printf(myurl)

	//prasing
	// result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	// qparams := result.Query()
	// fmt.Printf("Type of qparams is %T", qparams)

	// fmt.Println(qparams)

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=pankaj",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
