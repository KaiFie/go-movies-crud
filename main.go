package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
			
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "236915", Title: "Измены", Director: &Director{Firstname: "Евгений", Lastname: "Сладкоедов"}})
	movies = append(movies, Movie{ID: "2", Isbn: "281330", Title: "Модель XL", Director: &Director{Firstname: "Евгений", Lastname: "Лисицын"}})
	movies = append(movies, Movie{ID: "3", Isbn: "351237", Title: "Пацанки", Director: &Director{Firstname: "Евгений", Lastname: "Бюрбери"}})
	movies = append(movies, Movie{ID: "4", Isbn: "551034", Title: "Дорогая я забил", Director: &Director{Firstname: "Федор", Lastname: "Самойлов"}})
	movies = append(movies, Movie{ID: "5", Isbn: "891732", Title: "Немножко разведены", Director: &Director{Firstname: "Евгений", Lastname: "Виноградов"}})
	movies = append(movies, Movie{ID: "6", Isbn: "171304", Title: "Мужское и женское", Director: &Director{Firstname: "Эльдар", Lastname: "Приятный"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("PUT")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
