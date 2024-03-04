package main

import (
	"fmt"
	"log"
	"net/http"
   	"github.com/gorilla/mux"
)


func main()  {
	fmt.Println("Welcome to Mod in golang")	
	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":8080",r)

	log.Fatal(http.ListenAndServe(":8080",r))
}


func greeter()  {
	fmt.Println("Hey there mod user")
	
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Welcome to Golang Series by chandan"))
}