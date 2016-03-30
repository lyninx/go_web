package main

import (
	"fmt"
	"os"
	"html/template"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
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
	body := []byte("hi, world")	
	return &Page{Title: title, Body: body}, nil
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