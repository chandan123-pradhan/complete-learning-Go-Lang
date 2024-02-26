package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json="cource_name"`
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Welcome to the creat json ")
	EncodeJson()
}

func EncodeJson() {
	lcoCources := []course{
		{
			"React Js Bootcamp", 299, "Learncodeonline.in", "123456", []string{"web-dev", "js"},
		},
		{
			"Flutter Bootcamp", 199, "Learncodeonline.in", "12345", []string{"web-dev", "js"},
		},
		{
			"Firebase Bootcamp", 499, "Learncodeonline.in", "123456", []string{"web-dev", "js"},
		},
		{
			"xyz Bootcamp", 799, "Learncodeonline.in", "123456", []string{"web-dev", "js"},
		},
	}

	

	// finalJson, error := json.Marshal(lcoCources)  //encodeing json
	// if(error!=nil){
	// 	panic(error)
	// }

	finalJson, error := json.MarshalIndent(lcoCources,"","\t")  //encodeing json in json formatge
	if(error!=nil){
		panic(error)
	}

	fmt.Printf("%s\n",finalJson)  //printing that json
}
