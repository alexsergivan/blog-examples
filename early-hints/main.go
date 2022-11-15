package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed index.html
var index string

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/1", noHintsHandler)
	http.HandleFunc("/2", withHintsHandler)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}

func noHintsHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("")
	t, err := t.Parse(index)
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func withHintsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Link", "<https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/css/bootstrap.min.css>; rel=preload; as=style")
	w.Header().Add("Link", "<https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/js/bootstrap.bundle.min.js>; rel=preload; as=script")
	w.WriteHeader(http.StatusEarlyHints)
	t := template.New("")
	t, err := t.Parse(index)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, nil)
}
