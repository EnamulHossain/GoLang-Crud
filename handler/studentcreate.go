package handler

import (
	"StudentManagement/storage"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/justinas/nosurf"
)

// type StudentList struct{
// 	Students []storage.Student `db: "students"`
// }



func pareseStudentTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/createstudent.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err := t.ExecuteTemplate(w, "createstudent.html", data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) CreateStudent(w http.ResponseWriter, r *http.Request) {
	pareseStudentTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (c *connection) StoreStudent(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	form :=UserForm{}
	students := storage.Student{}

	if err := c.formDecoder.Decode(&students, r.PostForm); err != nil {
		log.Fatal(err)
	}

	form.Student = students
	if err := students.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		pareseStudentTemplate(w, form)
		return
	}

	_,err:= c.storage.CreateStudent(students)
	if err != nil {
		log.Fatalln(err)
	}

	http.Redirect(w, r, "/list/student", http.StatusSeeOther)
}
