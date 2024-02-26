package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com:3000/watch?v=cl7_ouTMFh0"

func main() {
	fmt.Println("Welcome to the Urls")

	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result) //for get all url.

	fmt.Println("scheme is =", result.Scheme) //output:- https

	fmt.Println("host is =", result.Host)  //for get host

	fmt.Println("port is=", result.Port())  //for get port number

	fmt.Println("raw query", result.RawQuery)//for get params

	queryParams :=result.Query();  // this will assign all params from url into queryParams map.

	fmt.Println("query params=",queryParams)
	fmt.Println("query params value is =",queryParams["v"])


	//for fetch all params one by one.

	for _, val := range queryParams{
		fmt.Println(val)
	}
}
