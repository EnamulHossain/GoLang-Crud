package handler

import (
	// "fmt"
	"StudentManagement/storage"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/justinas/nosurf"
)

func pareseEditUserTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/edituser.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "edituser.html", data)
}

func (c connection) EditUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	UserEdit,_:= c.storage.GetUserByID(id)
	
	var form UserForm
	form.User = *UserEdit
	form.CSRFToken = nosurf.Token(r)
	pareseEditUserTemplate(w, form)

}

func (c connection) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatalln("%#V", err)
	}

	var form UserForm
	user := storage.User{ID: uID}
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

	http.Redirect(w, r, "/user/list", http.StatusSeeOther)

}
