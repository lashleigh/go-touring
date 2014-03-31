package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

var (
	templates *template.Template
	places    []*Place
)

type Index struct {
	Places    []*Place
	PlacesStr string
	Host      string
}

func indexHandler(c http.ResponseWriter, req *http.Request) {
	str, err := json.Marshal(places)
	if err != nil {
		panic(err)
	}
	index := Index{places, string(str), req.Host}
	templates.Execute(c, index)
}

func main() {
	templates = template.Must(template.ParseFiles("index.html"))
	places = get_places(places)

	go h.run()
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/ws", wsHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
