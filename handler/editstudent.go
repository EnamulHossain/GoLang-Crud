package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/justinas/nosurf"
	// "github.com/go-chi/chi/v5"
)



func (c connection) StudentEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	editStudent,_:= c.storage.GetStudentByID(id)
	var form UserForm
	form.Student = *editStudent
	form.CSRFToken = nosurf.Token(r)
	c.pareseEditStudentTemplate(w, form)

}

func (h connection) StudentUpdate(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatalln("%#V", err)
	}

	var form UserForm


	
	student := storage.Student{ID: uID}
	if err := h.decoder.Decode(&student, r.PostForm); err != nil {
		log.Fatal(err)
	}
    form.Student = student
	if err := student.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.pareseStudentTemplate(w, form)
		return
	}

	
	_, eRr := h.storage.UpdateStudent(student)
	if eRr != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, ("/list/student"), http.StatusSeeOther)

}
