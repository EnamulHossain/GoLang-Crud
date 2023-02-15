package handler

import (
	"StudentManagement/storage"
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/justinas/nosurf"
)

//	type UserList struct {
//		Users []User `db:"users"`
//	}
type UserForm struct {
	User      storage.User
	Student   storage.Student
	Class     storage.Class
	Subject   storage.Subject
	ClassList []storage.Class
	FormError map[string]error
	CSRFToken string
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
			// fmt.Println(vErr)
			form.FormError = vErr
		}
		c.pareseRegTemplate(w, form)
		return
	}

	_, err := c.storage.CreateUser(user)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
