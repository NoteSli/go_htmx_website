package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, nil)
	}
	fs := http.FileServer(http.Dir("./css"))

	http.HandleFunc("/", h1)
	http.Handle("/css/", http.StripPrefix("/css", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
