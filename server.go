package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"strings"
	"html/template"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Page struct {
	URL string `json:"url"`
	Title string `json:"title"`
	Content string `json:"content"`
	Modified time.Time `json:"modified"`
}

type Pages []Page

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

////////////////////////////////////////////////////////////////////
var templatesPath = "templates/"
var apiPath = "api/"


func dbConnect() (*mgo.Session) {
	// connect to mongoDB
	session, err := mgo.Dial("mongodb://admin:testpass@ds023108.mlab.com:23108/go")
	if err != nil {
			panic(err)
			return nil
	}

	return session
}
func loadPage(url string) (*Page, error) {

	var session = dbConnect()
	pages := session.DB("go").C("pages")
	url = url[1:] //remove leading "/"
	result := Page{URL: url, Title: "not found", Content: "default page content"} //default page object

	if(url[:4] == apiPath){
		fmt.Print("API HIT ")
		url = strings.Replace(url, apiPath, "", 1)
	}

	fmt.Printf("====\n%-9s: /%s\n","url", url)
	var err = pages.Find(bson.M{"url": url}).One(&result)

	if err != nil {
		fmt.Printf("%-9s: %s\n", "content", "page not found")
		return &result, err
	} else {
		fmt.Printf("%-9s: %s\n", "title",result.Title)
		fmt.Printf("%-9s: %s\n ", "content",result.Content)	
	}
	
	return &result, nil
}

var templates = template.Must(template.ParseFiles(templatesPath+"index.template"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".template", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("template error")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r.URL.Path)
	
	if err != nil {
		// return if error
		//return
	}
	renderTemplate(w, "index", p)
}

func pageShow(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// pageId := vars["id"]
	p, err := loadPage(r.URL.Path)
	if err != nil {
		// return if error
		//return
	}
	renderTemplate(w, "index", p)
}

func apiPage(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r.URL.Path)
	if err != nil {
		// return if error
		fmt.Fprintf(w, "not found") 
		return
	}
	json.NewEncoder(w).Encode(p)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	// set static files path
	fs := justFilesFilesystem{http.Dir("./public/")}
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(fs))) 

	router.HandleFunc("/", Index)
	router.HandleFunc("/{id}", pageShow)
	router.HandleFunc("/"+ apiPath + "{id}", apiPage)

	fmt.Println("listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}