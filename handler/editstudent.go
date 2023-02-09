package handler

import (
	"StudentManagement/storage"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/justinas/nosurf"
	// "github.com/go-chi/chi/v5"
)

func pareseEditStudentTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/editstudent.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "editstudent.html", data)
}

func (c connection) StudentEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	editStudent,_:= c.storage.GetStudentByID(id)
	var form UserForm
	form.Student = *editStudent
	form.CSRFToken = nosurf.Token(r)
	pareseEditStudentTemplate(w, form)

}

func (h connection) StudentUpdate(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	// if err := r.ParseForm(); err != nil {
	// 	log.Fatalln("%#V", err)
	// }
	// //............................................................
	var form UserForm
	
	student := storage.Student{ID: uID}
	if err := h.formDecoder.Decode(&student, r.PostForm); err != nil {
		log.Fatal(err)
	}
    form.Student = student
	if err := student.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		pareseStudentTemplate(w, form)
		return
	}

	
	UpdateStudent,err := h.storage.UpdateStudent(student)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	// http.Redirect(w, r, ("/list/student"), http.StatusSeeOther)
	http.Redirect(w, r, fmt.Sprintln("/list/student",UpdateStudent), http.StatusSeeOther)
}
