package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func index(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	
	if err != nil {
		// return if error
		//return
	}
	renderTemplate(w, "index", p)
}

func page(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// pageId := vars["id"]
	p, err := loadPage(r)
	if err != nil {
		// return if error
		//return
	}
	renderTemplate(w, "index", p)
}

func create(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		// return if error
		//return
	}
	renderTemplate(w, "create", p)
}

func apiPage(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		// return if error
		fmt.Fprintf(w, "not found") 
		return
	}
	json.NewEncoder(w).Encode(p)
}