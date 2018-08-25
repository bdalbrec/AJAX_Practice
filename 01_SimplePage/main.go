package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/check", checker)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func checker(w http.ResponseWriter, req *http.Request) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Received %v\n", string(bs))

	io.WriteString(w, "We received "+string(bs))
}
