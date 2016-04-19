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
	renderTemplate(w, "page", p)
}

func create(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		// return if error
		//return
	}
	renderTemplate(w, "create", p)
}
func apiIndex(w http.ResponseWriter, r *http.Request) {
	p, err := loadPageList()
	if err != nil {
		// return if error
		fmt.Fprintf(w, "error") 
		return
	}
	json.NewEncoder(w).Encode(p)
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

func apiCreate(w http.ResponseWriter, r *http.Request) {
	err := createPage(r)
	
	if err != nil {
		fmt.Println(err)
	}
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	err := deletePage(r)
	if err != nil {
		fmt.Println(err)
	}	
}