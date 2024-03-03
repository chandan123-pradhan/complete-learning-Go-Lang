package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// models for cources
type Cource struct {
	CourceId    string  `json:"cource_id"`
	CourceName  string  `json:"cource_name"`
	CourcePrice int     `json:"cource_price"`
	Auther      *Auther `json:"auther"`
}

type Auther struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var cources []Cource

// Middlwere or helper method

func (c *Cource) IsEmpty() bool {
	return c.CourceName == ""
}

func main() {
	r := mux.NewRouter()

	//Seeding
	cources = append(cources, Cource{CourceId: "2", CourceName: "ReactJS", CourcePrice: 299, Auther: &Auther{FullName: "chandan pradhan", Website: "www.google.com"}})
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/get_cources", getAllCources).Methods("GET")
	cources = append(cources, Cource{CourceId: "3", CourceName: "flutter ", CourcePrice: 199, Auther: &Auther{FullName: "chandan", Website: "www.flutter.dev.com"}})
	
	

	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/get_allcources", getAllCources).Methods("GET")
	r.HandleFunc("/cource/{id}", getOneCource).Methods("GET")
	r.HandleFunc("/cource", createOneCource).Methods("POST")
	r.HandleFunc("/cource/{id}", updateOneCource).Methods("PUT")
	r.HandleFunc("/cource/{id}", deleteOneCource).Methods("DELETE")


	
	// listen to a port
	log.Fatal(http.ListenAndServe(":8080",r))
	
	
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Golang Series by chandan"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
	return
}

func getOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one cource")
	w.Header().Set("Content-Type", "application/json")

	//grab id from a request
	params := mux.Vars(r)

	//loop through cource and find maching cource id.
	for _, cource := range cources {
		if cource.CourceId == params["id"] {
			json.NewEncoder(w).Encode(cource)
			return
		}
	}
	json.NewEncoder(w).Encode("Not Found any cource for this id")
	return
}

func createOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one cource")
	w.Header().Set("Content-Type", "application/json")
	// what if entire body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	//what about - {}
	var cource Cource
	_ = json.NewDecoder(r.Body).Decode(&cource)

	if cource.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the json which you are sending")
		return
	}
	//generate a uniq id, string
	// apned cource into cources

	rand.Seed(time.Now().UnixNano())
	cource.CourceId = strconv.Itoa(rand.Intn(100))

	cources = append(cources, cource)

	json.NewEncoder(w).Encode("cources added successfully Done")

}

func updateOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one cource")
	w.Header().Set("Content-Type", "application/json")

	// first we need to grab id from request
	params := mux.Vars(r)

	// loop, id , remove , add with my ID

	for index, cource := range cources {
		if cource.CourceId == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)
			var cource Cource
			_ = json.NewDecoder(r.Body).Decode(&cource)
			cource.CourceId = params["id"]
			cources = append(cources, cource)
			json.NewEncoder(w).Encode(cource)
			return

		}

	}
	// TODO: Cource not found

}

func deleteOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one cource")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop, id, remove (index, index+1)

	for index, cource := range cources {
		if cource.CourceId == params["id"] {

			cources = append(cources[:index], cources[index+1:]...)
			json.NewEncoder(w).Encode("cources deleted successfully Done")

			break

		}

	}
	// TODO: Course not found
}
