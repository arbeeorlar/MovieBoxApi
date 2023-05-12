package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"ISBN"`
	Title    string    `json:"Title"`
	Director *Director `json:"director"`
}

type Director struct {
	firstName string `json:"firstname"`
	lastName  string `json:"lastname"`
}

var movies = []Movie{
	{ID: "1", Title: "Emi lokan", ISBN: "34555", Director: &Director{firstName: "Peter", lastName: "Daniel"}},
	{ID: "2", Title: "Iwo lokan", ISBN: "6758855", Director: &Director{firstName: "Tunde", lastName: "Fama"}},
	{ID: "3", Title: "Obi dient", ISBN: "980555", Director: &Director{firstName: "Tester", lastName: "Tester"}},
	{ID: "4", Title: "Ati kulate", ISBN: "3444455", Director: &Director{firstName: "Test", lastName: "Testers"}},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Method("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Method("GET")
	router.HandleFunc("/movies", createMovie).Method("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var item Movie
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	movies = append(movies, item)
	err = json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var item Movie
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	for index, newItem := range movies {
		if item.ID == newItem.ID {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	movies = append(movies, item)
	err = json.NewEncoder(w).Encode(&item)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
