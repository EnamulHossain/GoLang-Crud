package handler

import (
	"StudentManagement/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"golang.org/x/crypto/bcrypt"
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



func (c connection) Login(w http.ResponseWriter, r *http.Request) {
	c.pareseLoginTemplate(w, LoginUser{
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
			c.pareseLoginTemplate(w, lf)
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
			c.pareseLoginTemplate(w, lf)
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
		c.pareseLoginTemplate(w, lf)
		return
	}


	c.sessionManager.Put(r.Context(), "userName", (user.Name))
	http.Redirect(w, r, "/student/list", http.StatusSeeOther)
}
