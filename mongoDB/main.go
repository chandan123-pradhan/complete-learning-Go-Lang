package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandanpradhan-123/golang-project/routers"
)

func main() {
	fmt.Println("welcome to the mondo db")
	r := routers.Routers()
	fmt.Println("Server is getting started ...")

	log.Fatal(http.ListenAndServe(":5000", r))

	fmt.Println("listening port is 4000")

}
