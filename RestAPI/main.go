package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(rand.Intn(100000000))
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			var item Item
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.ID = strconv.Itoa(rand.Intn(100000000))
			items = append(items, item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(items)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)

}

func main() {
	// fmt.Println("Hello ")
	r := mux.NewRouter()

	// Mock Data
	items = append(items, Item{ID: "1", Isbn: "1234", Title: "Book One", Author: &Author{Firstname: "Lorem", Lastname: "Ipsome"}})

	r.HandleFunc("/api/items", getItems).Methods("GET")
	r.HandleFunc("/api/getitem/{id}", getItemByID).Methods("GET")
	r.HandleFunc("/api/additem", createItem).Methods("POST")
	r.HandleFunc("/api/updateitem/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/api/deleteitem/{id}", deleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
