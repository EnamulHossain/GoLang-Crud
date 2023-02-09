package handler

import (
	"html/template"
	"log"
	"net/http"
)

func pareseStudentListTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/studentlist.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "studentlist.html", data)
}

func (c connection) ListStudent(w http.ResponseWriter, r *http.Request) {
	
	listStudent,err:=c.storage.ListStudent()

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	pareseStudentListTemplate(w, listStudent)
}
