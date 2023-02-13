package handler

import (
	"StudentManagement/storage/postgres"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Name      string `form:"name"`
	Password  string `form:"password"`
	FormError map[string]error
	CSRFToken string
}

func (lu LoginUser) Validate() error {
	return validation.ValidateStruct(&lu,
		validation.Field(&lu.Name,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&lu.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

func pareseLoginTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/login.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "login.html", data)
}

func (c connection) Login(w http.ResponseWriter, r *http.Request) {
	pareseLoginTemplate(w, LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}

func (c connection) LoginPostHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	var lf LoginUser

	if err := c.decoder.Decode(&lf, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if err := lf.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			formErr := make(map[string]error)
			for key, val := range vErr {
				formErr[strings.Title(key)] = val
			}
			lf.FormError = formErr
			lf.Password = ""
			lf.CSRFToken = nosurf.Token(r)
			pareseLoginTemplate(w, lf)
			return
		}
	}

	user, err := c.storage.GetUserByUsername(lf.Name)
	if err != nil {
		if err.Error() == postgres.NotFound {
			formErr := make(map[string]error)
			formErr["Name"] = fmt.Errorf("credentials does not match")
			lf.FormError = formErr
			lf.CSRFToken = nosurf.Token(r)
			lf.Password = ""
			pareseLoginTemplate(w, lf)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}


	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lf.Password)); err != nil {
		formErr := make(map[string]error)
		formErr["Name"] = fmt.Errorf("credentials does not match")
		lf.FormError = formErr
		lf.CSRFToken = nosurf.Token(r)
		lf.Password = ""
		pareseLoginTemplate(w, lf)
		return
	}


	c.sessionManager.Put(r.Context(), "userName", (user.Name))
	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
