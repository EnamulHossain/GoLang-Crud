package handler

import (
	"html/template"
	"log"
	"net/http"
)

func pareseHomeTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/home.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "home.html", data)
}

func (h connection) Home(w http.ResponseWriter, r *http.Request) {
	pareseHomeTemplate(w, r)
}