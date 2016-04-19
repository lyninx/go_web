package main

import (
	"time"
	"strings"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func loadPage(r *http.Request) (*Page, error) {
	var url = r.URL.Path
	var start = time.Now()
	
	pages := session.DB("go").C("pages")

	url = url[1:] //remove leading "/"
	result := Page{URL: "", Title: "not found", Content: "sorry, couldn't find that page."} //default page object

	// remove leading api path
	url = strings.Replace(url, apiPath, "", 1)
	
	var err = pages.Find(bson.M{"url": url}).One(&result)

	logRequest(r, time.Since(start))

	if err != nil {
		return &result, err
	} 

	return &result, nil
}

func createPage(r *http.Request) (error) {
	var start = time.Now()
	decoder := json.NewDecoder(r.Body)
	p := Page{Modified: start} 
	err := decoder.Decode(&p)

	logRequest(r, time.Since(start))

	if err != nil {
		return err
	}

	pages := session.DB("go").C("pages")
	err = pages.Insert(&p)

	if err != nil {
		return err
	}
	return nil
}
func deletePage(r *http.Request) (error) {
	var url = r.URL.Path
	var start = time.Now()
	
	pages := session.DB("go").C("pages")

	url = url[1:] //remove leading "/"	
	url = strings.Replace(url, apiPath, "", 1) // remove leading api path
	
	var err = pages.Remove(bson.M{"url": url})

	logRequest(r, time.Since(start))

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