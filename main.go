package main

import (
	"fmt"
	"html/template"
	"log"

	"net/http"
)

func main() {
	fmt.Println("HELLO WORLD")

	//Basic example to show in web
	// h1 := func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "HEY\n")
	// 	io.WriteString(w, r.Method)
	// }

	//Parse the html and show in browser, data driven templates
	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
	}

	http.HandleFunc("/", h1)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Problem starting server", err)
	}

}
