package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type type_page struct {
	title string
	body  []byte
}

func save_page(p *type_page) error {
	filename := p.title + ".txt"
	return os.WriteFile(filename, p.body, 0600)
}

func load_page(page_title string) (*type_page, error) {
	filename := page_title + ".txt"
	page_body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &type_page{title: page_title, body: page_body}, nil
}

func view_handler(response http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	p, err := load_page(title)
	if err != nil {
		http.Redirect(response, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(response, "view", p)
}

func save_handler(response http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/save/"):]
	body := request.FormValue("body")
	p := &type_page{title: title, body: []byte(body)}
	err := save_page(p)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, "/view/"+title, http.StatusFound)
}

func edit_handler(response http.ResponseWriter, request *http.Request) {
	page_title := request.URL.Path[len("/edit/"):]
	p, err := load_page(page_title)
	if err != nil {
		p = &type_page{title: page_title}
	}
	renderTemplate(response, "edit", p)
}

func renderTemplate(response http.ResponseWriter, tmpl string, p *type_page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(response, p)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/", view_handler)
	http.HandleFunc("/edit/", edit_handler)
	http.HandleFunc("/save/", save_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
