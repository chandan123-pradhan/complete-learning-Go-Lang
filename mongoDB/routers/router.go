package routers

import (
	"github.com/chandanpradhan-123/golang-project/controller"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	router := mux.NewRouter();

	router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}",controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/delete_all_movie",controller.DeleteAllMovie).Methods("DELETE")
	return router
}