package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Artist struct {
	Name string
	Year int
	Id   int
}

var artists = []Artist{
	{Name: "Queen", Year: 1970, Id: 1},
	{Name: "Daft Punk", Year: 1993, Id: 2},
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("Id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "400 - ID invalide", http.StatusBadRequest)
		return
	}

	var found *Artist
	for i := range artists {
		if artists[i].Id == id {
			found = &artists[i]
			break
		}
	}

	if found == nil {
		http.Error(w, "404 - Artiste introuvable", http.StatusNotFound)
		return
	}

	tpl := template.Must(template.ParseFiles("templates/artist.html"))
	tpl.Execute(w, *found)
}

func list(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/artists.html"))
	tpl.Execute(w, artists)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour le web !")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/artists/", list)
	http.HandleFunc("/artists/{Id}", artistPage)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}