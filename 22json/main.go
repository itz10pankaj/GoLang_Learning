package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`              // ye field visible nhi hogi
	Tags     []string `json:"tags,omitempty"` // if any field is given to be nil then omitempty will not allow that one to appear
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "learcncode.in", "o12345678", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "learcncode.in", "o123bfhj5678", []string{"web-dev", "Mern"}},
		{"Angular Bootcamp", 399, "learcncode.in", "hj5678", nil},
	}

	//package this datra as json
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	fmt.Printf("%T", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "Reactjs Bootcamp",
                "Price": 299,
                "Platform": "learcncode.in",
                "Tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json in valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON in not valid")
	}

	//SOME CASES WHERE YOU JUST WANT TO ADD DATA TO KEY VALUE
	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	// fmt.Printf("%#v\n", myonlineData)
	for k, v := range myonlineData {
		fmt.Printf("key is %v and value is %v", k, v)
	}
}
