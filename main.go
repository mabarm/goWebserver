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

	//For record
	type Film struct {
		Title    string
		Director string
	}

	//Parse the html and show in browser, data driven templates
	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "Blade", Director: "Ridley"},
				{Title: "Godfather", Director: "Francis"},
			},
		}
		//films is the data we want to inject in html dynamically, ultimately html will display this.
		temp.Execute(w, films)
	}

	http.HandleFunc("/", h1)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Problem starting server", err)
	}

}
