package handler

import (
	"html/template"
	"log"
	"net/http"
)

func pareseUserTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/userlist.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "userlist.html", data)
}

func (c connection) UserList(w http.ResponseWriter, r *http.Request) {
	var user []User
	if err := c.db.Select(&user, "SELECT * FROM users"); err != nil {
		log.Fatal(err)
	}
	pareseUserTemplate(w, user)
}
