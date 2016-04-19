package main

import (
	"net/http"
	"encoding/json"
)

func index(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		//logError(r, err)
	}
	renderTemplate(w, "index", p)
}

func page(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		//logError(r, err)
		renderTemplate(w, "error", p)
		return
	}
	renderTemplate(w, "page", p)
}

func create(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		//logError(r, err)
	}
	renderTemplate(w, "create", p)
}
func apiIndex(w http.ResponseWriter, r *http.Request) {
	p, err := loadPageList()
	if err != nil {
		logError(r, err)
		return
	}
	json.NewEncoder(w).Encode(p)
}
func apiPage(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r)
	if err != nil {
		logError(r, err)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func apiCreate(w http.ResponseWriter, r *http.Request) {
	err := createPage(r)
	if err != nil {
		logError(r, err)
	}
}

func apiDelete(w http.ResponseWriter, r *http.Request) {
	err := deletePage(r)
	if err != nil {
		logError(r, err)
	}	
}