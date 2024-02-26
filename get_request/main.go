package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the get request")
	// getRequest()
	// postRequest();
	postRequestWithFormBody()

}

func getRequest() {
	url := "https://dummy.restapiexample.com/api/v1/employees"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code is ", response.StatusCode)

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	//fmt.Println("data is=", string(data) )// by this line we can get all data.
	convertIntoByteStringAndDecodeData(data)
	

}

func postRequest() {
	url := "https://dummy.restapiexample.com/api/v1/create"


	requestBody :=strings.NewReader(`{
		"name":"test",
		"salary":"123",
		"age":"23"
	}`);


	response, err := http.Post(url,"application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code is ", response.StatusCode)

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	//fmt.Println("data is=", string(data) )// by this line we can get all data.
	convertIntoByteStringAndDecodeData(data)
	

}


func postRequestWithFormBody() {
	myUrl := "https://dummy.restapiexample.com/api/v1/create"


	requestBody :=url.Values{};

	requestBody.Add("name","test",)
	requestBody.Add("salary","123",)
	requestBody.Add("age","23")
	
	response, err := http.PostForm(myUrl, requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code is ", response.StatusCode)

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	//fmt.Println("data is=", string(data) )// by this line we can get all data.
	convertIntoByteStringAndDecodeData(data)
	

}


func convertIntoByteStringAndDecodeData(data []byte)  {
	var responseString strings.Builder

	byteCount, _ := responseString.Write(data)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}
