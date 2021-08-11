package main

import (
	"log"
	"net/http"
	"rtforum/config"
	"text/template"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseGlob("web/index.html")
	if err != nil {
		log.Fatalf("Can't load index template: %s", err)
	}

	web := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", web))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(config.CLIENT_PORT, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
