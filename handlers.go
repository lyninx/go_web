package main

import (
	"fmt"
	"time"
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

	decoder := json.NewDecoder(r.Body)
	p := Page{Modified: time.Now()} 
	err := decoder.Decode(&p)

	if err != nil {
		// return if error
		fmt.Fprintf(w, "new page") 
		return
	}

	var session = dbConnect()
	pages := session.DB("go").C("pages")
	err = pages.Insert(&p)

	if err != nil {
		fmt.Println(err)
	}
}