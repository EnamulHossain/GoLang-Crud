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
	// "github.com/go-chi/chi/v5"
)



func (c connection) ClassEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	editClass,_:= c.storage.GetClassByID(id)
	var form UserForm
	form.Class = *editClass
	form.CSRFToken = nosurf.Token(r)
	c.pareseClassEditTemplate(w, form)

}

func (h connection) ClassUpdate(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatalln("%#V", err)
	}

	var form UserForm
	class := storage.Class{ID: uID}
	if err := h.decoder.Decode(&class, r.PostForm); err != nil {
		log.Fatal(err)
	}
    form.Class = class
	
	if err := class.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.pareseClassTemplate(w, form)
		return
	}

	
	_, eRr := h.storage.UpdateClass(class)
	if eRr != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, ("/class/list"), http.StatusSeeOther)

}
