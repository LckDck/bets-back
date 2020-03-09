package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getItems(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/items" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		value, _ := json.Marshal("Списюля!")
		w.Write(value)
	}
}

func main() {
	http.HandleFunc("/items", getItems)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
