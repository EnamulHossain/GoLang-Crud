package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"

	// validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

func (c connection) CreateSubject(w http.ResponseWriter, r *http.Request) {
	classList, err := c.storage.ListClass()
	if err != nil {
		log.Fatalln(err)
	}

	c.pareseSubjectTemplate(w, UserForm{
		ClassList: classList,
		CSRFToken: nosurf.Token(r),
	})
}

func (c *connection) StoreSubject(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// form := UserForm{}
	subjects := storage.Subject{}

	if err := c.decoder.Decode(&subjects, r.PostForm); err != nil {
		log.Fatal(err)
	}


	_, err := c.storage.CreateSubject(subjects)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/subject/list", http.StatusSeeOther)
}
