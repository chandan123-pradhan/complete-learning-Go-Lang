package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"cource_name"` //here we are giving custon key name for use in json.
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"password"`
	Tags     []string `json:"tags"`
}

func main() {
	fmt.Println("Welcome to the creat json ")
	//EncodeJson()
	DecodeJson()
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

	finalJson, error := json.MarshalIndent(lcoCources, "", "\t") //encodeing json in json formatge
	if error != nil {
		panic(error)
	}

	fmt.Printf("%s\n", finalJson) //printing that json
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"Name": "React Js Bootcamp",
		"Price": 299,
		"Platform": "Learncodeonline.in",
		"Password": "123456",
		"Tags": ["web-dev","js"]
	}
`)
	var lcoCource course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		//it means json data is valid
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCource) //encoding json.
		fmt.Printf("%#v\n", lcoCource)
	} else {
		fmt.Println("json was not valid")
	}

	// some case where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("key is %v and value is %v\n",k,v)
	}

}
