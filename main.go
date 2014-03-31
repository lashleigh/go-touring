package main

import (
	"log"
	"net/http"
	"text/template"
)

var (
	templates *template.Template
)

func indexHandler(c http.ResponseWriter, req *http.Request) {
	log.Println(c, req, req.Host)
	templates.Execute(c, req.Host)
}

func main() {
	templates = template.Must(template.ParseFiles("index.html", "_place.html"))

	go h.run()
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/ws", wsHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
