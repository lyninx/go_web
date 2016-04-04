package main

import (
	"fmt"
	"os"
	"log"
	"html/template"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	//"encoding/json"
)

type Page struct {
	URL string `json:"url"`
	Title string `json:"title"`
	Content string `json:"content"`
}

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

func loadPage(url string) (*Page, error) {
	// connect to mongoDB
	session, err := mgo.Dial("mongodb://admin:testpass@ds023108.mlab.com:23108/go")
		if err != nil {
				panic(err)
		}
		defer session.Close()

	pages := session.DB("go").C("pages")


	// page := &Page{Title: "test", Content: "pagecontent"}
	// b, err := json.Marshal(page)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))
	// err = pages.Insert(page)
	// if err != nil {
	// 		log.Fatal(err)
	// }

	result := Page{URL: url, Title: "not found", Content: "default page content"}
	url = url[1:]
	fmt.Printf("====\n%-9s: /%s\n","url", url)
	err = pages.Find(bson.M{"url": url}).One(&result)
	if err != nil {
		fmt.Printf("%-9s: %s\n", "content","page not found")
	} else {
		fmt.Printf("%-9s: %s\n", "title",result.Title)
		fmt.Printf("%-9s: %s\n ", "content",result.Content)	
	}
	
	return &result, nil
}

var templatesDir = "templates/"
var templates = template.Must(template.ParseFiles(templatesDir+"index.template"))

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
		// redirect if error
		return
	}
	renderTemplate(w, "index", p)
}

func pageShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pageId := vars["id"]
    p, err := loadPage(r.URL.Path)
    fmt.Println(pageId)
	if err != nil {
		// redirect if error
		return
	}
    renderTemplate(w, "index", p)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	// set static files path
	fs := justFilesFilesystem{http.Dir("./public/")}
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(fs))) 


	//http.HandleFunc("/", makeHandler(indexHandler))
	//router.HandleFunc("/", makeHandler(indexHandler))
	router.HandleFunc("/", Index)
	router.HandleFunc("/{id}", pageShow)
	//router.PathPrefix("/public/").Handler(http.FileServer(fs))
    // http.HandleFunc("/posts", postIndex)
    // http.HandleFunc("/posts/{postId}", postShow)
	fmt.Println("listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}