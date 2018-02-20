package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Name string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	// dynamically sets the name from the query parameter
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Abaddon"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}

		// returns an error object if present
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}