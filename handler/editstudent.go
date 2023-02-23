package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

func (c connection) StudentEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	editStudent, _ := c.storage.GetStudentByID(id)

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

	students := storage.Student{ID: uID}
	
	if err := h.decoder.Decode(&students, r.PostForm); err != nil {
		log.Fatal(err)
	}
	
	var form UserForm


	if err := students.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.pareseStudentTemplate(w, form)
		return
	}

	_, eRr := h.storage.UpdateStudent(students)
	if eRr != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, ("/student/list"), http.StatusSeeOther)

}
