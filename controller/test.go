package main

import (
	"net/http"
	"html/template"
)

func view(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("pathfinder_form.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", view)
	http.ListenAndServe(":8080", nil)
}
