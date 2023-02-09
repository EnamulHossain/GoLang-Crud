package handler

import (
	"StudentManagement/storage"
	"html/template"
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
	FormError map[string]error
	CSRFToken string
}

func pareseRegTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/reg.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err := t.ExecuteTemplate(w, "reg.html", data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) Reg(w http.ResponseWriter, r *http.Request) {
	pareseRegTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (c connection) StoreUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	form := UserForm{}
	user := storage.User{}

	if err := c.formDecoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}

	form.User = user
	
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			// fmt.Println(vErr)
			form.FormError = vErr
		}
		pareseRegTemplate(w, form)
		return
	}

	_, err := c.storage.CreateUser(user)
	if err != nil {
		log.Fatalln(err)
	}
	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
