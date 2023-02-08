package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	// "golang.org/x/crypto/bcrypt"
	// validation "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-ozzo/ozzo-validation/v4/is"
)

//	type login struct{
//		username string
//		password string
//	}
func pareseLoginTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/login.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "login.html", data)
}

func (c connection) Login(w http.ResponseWriter, r *http.Request) {
	pareseLoginTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (c *connection) LoginUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// Name := r.FormValue("username")
	// Password := r.FormValue("password")

	// l := login{
	// 	username: Name,
	// 	password: Password,
	// }

	// compareLogin:=

	http.Redirect(w, r, "/user/list", http.StatusSeeOther)

}

// // package handler

// import (
// 	"html/template"
// 	"log"
// 	"net/http"

// 	validation "github.com/go-ozzo/ozzo-validation/v4"
// 	"github.com/go-ozzo/ozzo-validation/v4/is"
// 	"github.com/gorilla/csrf"
// )

// type Login struct {
// 	Email    string
// 	Password string
// }

// type LoginTempData struct {
// 	CSRFField  template.HTML
// 	Form       Login
// 	FormErrors map[string]string
// }

// func (l Login) Validate() error {
// 	return validation.ValidateStruct(&l,
// 		validation.Field(&l.Email, validation.Required, is.Email),
// 		validation.Field(&l.Password, validation.Required, validation.Length(6, 12)),
// 	)
// }

// func (s *Server) getLogin(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Method: getLogin")
// 	formData := LoginTempData{
// 		CSRFField: csrf.TemplateField(r),
// 	}
// 	s.loadLoginTemplate(w, r, formData)
// }

// func (s *Server) postLogin(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Method: postLogin")

// 	if err := r.ParseForm(); err != nil {
// 		log.Fatalln("parsing error")
// 	}

// 	var form Login
// 	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
// 		log.Fatalln("decoding error")
// 	}

// 	if err := form.ValidateL(); err != nil {
// 		vErrs := map[string]string{}
// 		if e, ok := err.(validation.Errors); ok {
// 			if len(e) > 0 {
// 				for key, value := range e {
// 					vErrs[key] = value.Error()
// 				}
// 			}
// 		}

// 		data := LoginTempData{
// 			CSRFField:  csrf.TemplateField(r),
// 			Form:       form,
// 			FormErrors: vErrs,
// 		}
// 		s.loadLoginTemplate(w, r, data)
// 		return
// 	}

// 	hash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
// 	if err:= bcrypt.CompareHashAndPassword(hash, []byte("123456")); err != nil{
// 		log.Fatalf("Password does not match ")
// 	}

// 	session, _ := s.session.Get(r, "practice_session")
// 	session.Values["user_id"] = "1"
// 	if err := session.Save(r, w); err != nil {
// 		log.Fatalln("error while saving user id into session")
// 	}
// 	http.Redirect(w, r, "/student/create", http.StatusTemporaryRedirect)
// }

// func (s *Server) loadLoginTemplate(w http.ResponseWriter, r *http.Request, form LoginTempData) {
// 	tmp := s.templates.Lookup("login.html")
// 	if err := tmp.Execute(w, form); err != nil {
// 		log.Println("Error executing template :", err)
// 		return
// 	}
// }
