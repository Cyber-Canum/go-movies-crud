package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteFromMovies).Methods("DELETE")

	fmt.Printf("Starting server at port:8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, v := range movies {

		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			break
		}

	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Int())
	movies := append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	//DELETE
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteMovie(params["id"])

	//ADD
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	movies := append(movies, movie)

	//Print
	json.NewEncoder(w).Encode(movies)

}

func deleteFromMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(id string) {
	for i, v := range movies {
		if v.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			return
		}
	}
}

func init() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "54789",
		Title: "Forest Gump",
		Director: &Director{
			FirstName: "Frank",
			LastName:  "Drew",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "98712",
		Title: "May",
		Director: &Director{
			FirstName: "Yakinoe",
			LastName:  "Li",
		},
	})

	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "130009",
		Title: "Beast",
		Director: &Director{
			FirstName: "Paul",
			LastName:  "Shewnich",
		},
	})
}
