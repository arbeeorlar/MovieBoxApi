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
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie() {

}
func createMovie() {

}
func updateMovie() {

}
func deleteMovie() {

}
