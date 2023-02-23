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



func (c connection) EditSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	editSubject,_:= c.storage.GetSubjectByID(id)         //
	var form UserForm
	form.Subject = *editSubject
	form.CSRFToken = nosurf.Token(r)
	c.pareseEditSubjectTemplate(w, form)

}




func (h connection) SubjectUpdate(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatalln("%#V", err)
	}

	var form UserForm


	
	subject := storage.Subject{ID: uID}
	if err := h.decoder.Decode(&subject, r.PostForm); err != nil {
		log.Fatal(err)
	}


    form.Subject = subject
	if err := subject.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.pareseSubjectTemplate(w, form)
		return
	}

	
	_, eRr := h.storage.UpdateSubject(subject)
	if eRr != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, ("/subject/list"), http.StatusSeeOther)

}
