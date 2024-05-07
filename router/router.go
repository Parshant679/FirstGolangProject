package router

import (
	controller "Api/controllers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("this is home router")
	})
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.CreateMovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.CreateMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.CreateMovie).Methods("DELETE")

	return router
}
