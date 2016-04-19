package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"html/template"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
)

type Page struct {
	URL string `json:"url"`
	Title string `json:"title"`
	Content string `json:"content"`
	Modified time.Time `json:"modified"` 
}

type M map[string]interface{}

type Pages []Page

var templatesPath = "templates/"
var apiPath = "api/"

type justFilesFilesystem struct {
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func dbConnect() (*mgo.Session) {
	// get mongoDB session
	session, err := mgo.Dial("mongodb://admin:testpass@ds023108.mlab.com:23108/go")
	if err != nil {
			panic(err)
			return nil
	}

	return session
}
// connect to database
var session = dbConnect()

// load template files
var templates = template.Must(template.ParseFiles(
	templatesPath+"index.template",
	templatesPath+"page.template",
	templatesPath+"create.template",
	templatesPath+"header.template",
	templatesPath+"footer.template"))

//render page
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	index, err := loadPageList()
	pageData := M{
		"page": p,
		"index": index,
	}
	err = templates.ExecuteTemplate(w, "header.template", p)
	err = templates.ExecuteTemplate(w, tmpl+".template", pageData)
	err = templates.ExecuteTemplate(w, "footer.template", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("template error")
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// set static files path
	fs := justFilesFilesystem{http.Dir("./public/")}
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(fs))) 

	// call handler functions based on route
	router.HandleFunc("/"+ apiPath, apiIndex).Methods("GET")
	router.HandleFunc("/"+ apiPath + "create", apiCreate).Methods("POST")
	router.HandleFunc("/"+ apiPath + "{id}", apiDelete).Methods("DELETE")
	router.HandleFunc("/"+ apiPath + "{id}", apiPage).Methods("GET")
	router.HandleFunc("/", index)
	router.HandleFunc("/create", create)
	router.HandleFunc("/{id}", page)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}