package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)






func (c connection) CreateClass(w http.ResponseWriter, r *http.Request) {
	c.pareseClassTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (c *connection) StoreClass(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	form :=UserForm{}
	classes := storage.Class{}

	if err := c.decoder.Decode(&classes, r.PostForm); err != nil {
		log.Fatal(err)
	}


	form.Class = classes
	if err := classes.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		c.pareseClassTemplate(w, form)
		return
	}

	 c.storage.CreateClass(classes)
	

	http.Redirect(w, r, "/class/list", http.StatusSeeOther)
}
