package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang-crud/entity"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	movies = append(movies, entity.Movie{ID: "1", ISBN: "111", Title: "Movie One", Director: &entity.Director{
		Firstname: "Andre",
		Lastname:  "Rizaldi Brillianto",
	}})
	movies = append(movies, entity.Movie{ID: "2", ISBN: "222", Title: "Movie Two", Director: &entity.Director{
		Firstname: "Alice",
		Lastname:  "Aureliax",
	}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8181")
	log.Fatal(http.ListenAndServe(":8181", r))
}

var movies []entity.Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for key, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:key], movies[key+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, value := range movies {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie entity.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// merubah id menjadi antara 1 - 100000000
	//movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for key, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:key], movies[key+1:]...)
			var movie entity.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
