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
		templ := template.Must(template.ParseFiles("views/index.html"))
		templ.Execute(w, nil)
	}
	fs := http.FileServer(http.Dir("./css"))
	image_fs := http.FileServer(http.Dir("./images"))
	font_fs := http.FileServer(http.Dir("./fonts/font-awesome-4.7.0/css"))
	other_font_fs := http.FileServer(http.Dir("./fonts"))

	http.HandleFunc("/", h1)
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.Handle("/images/", http.StripPrefix("/images", image_fs))
	http.Handle("/fonts/font-awesome-4.7.0/css/", http.StripPrefix("/fonts/font-awesome-4.7.0/css", font_fs))
	http.Handle("/fonts/", http.StripPrefix("/fonts", other_font_fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
