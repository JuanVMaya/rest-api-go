package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Item struct {
	UID string `json:"UID"`
	Name string `json:"Name"`
	Description string `json:"Description"`
	Price float64 `json:"Price"`
}

var inventory []Item	

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range inventory {
		if item.UID == params["UID"] {
			//Delete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(inventory)
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getInventory")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var item Item
	_=json.NewDecoder(r.Body).Decode(&item)

	inventory=append(inventory, item)

	json.NewEncoder(w).Encode(item)
}



func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{UID}", deleteItem).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	inventory=append(inventory, Item{
		UID: "0", 
		Name: "Toyota RAV4 2019", 
		Description: "A 4-door SUV", 
		Price: 25000.0,
	})
	inventory=append(inventory, Item{
		UID: "1", 
		Name: "Toyota Camry 2019", 
		Description: "A 4-door car", 
		Price: 20000.0,
	})
	handleRequests()
}