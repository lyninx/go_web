package main

import (
	"log"
	"time"
	"strings"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func loadPage(r *http.Request) (*Page, error) {
	var url = r.URL.Path
	var start = time.Now()
	
	pages := session.DB("go").C("pages")

	url = url[1:] //remove leading "/"
	result := Page{URL: "", Title: "not found", Content: "default page content"} //default page object

	// remove leading api path
	url = strings.Replace(url, apiPath, "", 1)
	
	var err = pages.Find(bson.M{"url": url}).One(&result)

	log.Printf(
		"%-9s\t%-9s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	if err != nil {
		return &result, err
	} 

	return &result, nil
}

func deletePage(r *http.Request) (error) {
	var url = r.URL.Path
	var start = time.Now()
	
	pages := session.DB("go").C("pages")

	url = url[1:] //remove leading "/"	
	url = strings.Replace(url, apiPath, "", 1) // remove leading api path
	
	var err = pages.Remove(bson.M{"url": url})

	log.Printf(
		"%-9s\t%-9s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	if err != nil {
		return err
	} 

	return nil
}

func loadPageList() (*Pages, error){
	pages := session.DB("go").C("pages")
	result := Pages{}

	var err = pages.Find(nil).All(&result)
	if err != nil {
		return &result, err
	} 

	return &result, nil
}