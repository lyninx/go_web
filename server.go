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
)

type Page struct {
	URL string `json:"url"`
	Title string `json:"title"`
	Content string `json:"content"`
	Modified time.Time `json:"modified"` 
}

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
	// connect to mongoDB
	session, err := mgo.Dial("mongodb://admin:testpass@ds023108.mlab.com:23108/go")
	if err != nil {
			panic(err)
			return nil
	}

	return session
}
var session = dbConnect()

func loadPage(r *http.Request) (*Page, error) {
	var url = r.URL.Path
	var start = time.Now()
	
	pages := session.DB("go").C("pages")
	fmt.Println(url)

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

func loadPageList(r *http.Request) (*Pages, error){
	var start = time.Now()
	pages := session.DB("go").C("pages")

	result := Pages{}

	var err = pages.Find(nil).All(&result)

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

var templates = template.Must(template.ParseFiles(
	templatesPath+"index.template",
	templatesPath+"page.template",
	templatesPath+"create.template",
	templatesPath+"header.template",
	templatesPath+"footer.template"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, "header.template", p)
	err = templates.ExecuteTemplate(w, tmpl+".template", p)
	err = templates.ExecuteTemplate(w, "footer.template", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("template error")
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
	router.HandleFunc("/"+ apiPath + "{id}", apiPage).Methods("GET")
	router.HandleFunc("/", index)
	router.HandleFunc("/create", create)
	router.HandleFunc("/{id}", page)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}