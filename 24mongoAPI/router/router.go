package router

import (
	"mongoApis/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// fmt.Println("Router")
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllRecordsHandler).Methods("GET")
	router.HandleFunc("/api/create-movie", controller.CreateMovieHandler).Methods("POST")
	router.HandleFunc("/api/mark-movie-watched/{id}", controller.MarkAsWatchedHandler).Methods("PUT")
	router.HandleFunc("/api/delete-movie/{id}", controller.DeleteOneMovieHandler).Methods("DELETE")
	router.HandleFunc("/api/delete-all", controller.DeleteAllMoviesHandler).Methods("DELETE")

	return router
}
