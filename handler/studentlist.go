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
	var student []Student
	if err := c.db.Select(&student, "SELECT * FROM students WHERE deleted_at IS NULL ORDER BY id ASC"); err != nil {
		log.Fatal(err)
	}
	pareseStudentListTemplate(w, student)
}
