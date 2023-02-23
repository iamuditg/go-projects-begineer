package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	// add static data in movies
	staticData()

	router.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	router.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Starting server at port at 8000\n")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func staticData() {

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "43a227",
		Title: "Movie One",
		Director: &Director{
			FirstName: "Ud",
			LastName:  "it",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "43a227",
		Title: "Movie One",
		Director: &Director{
			FirstName: "Ud",
			LastName:  "it",
		},
	})

}
