package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Item Struct (Model)
type Item struct {
	ID     string  `json:"id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "Author"`
}

//  Author Struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

// Init Items struct as a slice Item struct
var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)

}

func getItemByID(w http.ResponseWriter, r *http.Request) {

}

func createItem(w http.ResponseWriter, r *http.Request) {

}

func updateItem(w http.ResponseWriter, r *http.Request) {

}

func deleteItem(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// fmt.Println("Hello ")
	r := mux.NewRouter()

	// Mock Data
	items = append(items, Item{ID: "1", Isbn: "1234", Title: "Book One", Author: &Author{Firstname: "Lorem", Lastname: "Ipsome"}})

	r.HandleFunc("/api/items", getItems).Methods("GET")
	r.HandleFunc("/api/getItems/{id}", getItemByID).Methods("GET")
	r.HandleFunc("/api/addItem", createItem).Methods("POST")
	r.HandleFunc("/api/updateItem", updateItem).Methods("PUT")
	r.HandleFunc("/api/item/deleteItem", deleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
