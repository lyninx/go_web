package main

import (
	"fmt"
	"os"
	"html/template"
	"net/http"
	"regexp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"encoding/json"
)

type Page struct {
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

func loadPage(title string) (*Page, error) {
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

	result := Page{Title: title, Content: "default page content"}
	title = title[1:]
	err = pages.Find(bson.M{"title": title}).One(&result)
	if err != nil {
			fmt.Println("page not found")
	}
	fmt.Println(title)
	fmt.Println("page:", result.Content)

	/////////////////////
	return &result, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	
	if err != nil {
		http.Redirect(w, r, "/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "index", p)
}

var templates = template.Must(template.ParseFiles("index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



var validPath = regexp.MustCompile("(/([a-zA-Z0-9]*))+")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			fmt.Println(r.URL.Path)
			return
		}
		fn(w, r, m[0])
	}
}



func main() {

	fs := justFilesFilesystem{http.Dir("public/")}
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(fs)))

	http.HandleFunc("/", makeHandler(indexHandler))
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}