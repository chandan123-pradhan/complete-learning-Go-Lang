package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to the webrequst")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is a type of %T\n", res)
	
	defer res.Body.Close() //callers responsibility to close the connections.

	dataByte , err :=io.ReadAll(res.Body)

	if(err!=nil){
		panic(err)
	}

	content :=string(dataByte);

	fmt.Println("content of the response body \n",content)
}
