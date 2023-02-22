package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

type MarkForm struct {
	Class           string
	Student         string
	CSRFToken      string
}


type UserForm struct {
	User           storage.User
	Student        storage.Student
	StudentList    []storage.Student
	Class          storage.Class
	Subject        storage.Subject
	SubjectList    []storage.Subject
	ClassList      []storage.Class
	StudentSubject []storage.StudentSubject
	MarkInput      []storage.MarkInputStore
	FormError      map[string]error
	CSRFToken      string
}

func (c connection) Reg(w http.ResponseWriter, r *http.Request) {
	c.pareseRegTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (c connection) StoreUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	form := UserForm{}
	user := storage.User{}

	if err := c.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}

	form.User = user

	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		c.pareseEditUserTemplate(w, form)
		return
	}

	_, err := c.storage.CreateUser(user)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
