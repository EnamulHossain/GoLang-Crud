package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	// "github.com/go-playground/form"
	"github.com/justinas/nosurf"
)

type User struct {
	ID        int       `db:"id" form:"-"`
	Name      string    `db:"name" form:"name"`
	Email     string    `db:"email" form:"email"`
	Password  string    `db:"password" form:"password"`
	CreatedAt time.Time `db:"created_at" form:"created_at"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at"`
	CSRFToken string    `db:"-" form:"csrf_token"`
	FormError map[string]error
}

func (u User) Validate() error {
	vre := validation.Required.Error
	len := validation.Length(2, 20).Error
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name,
			vre("The name  is required"),
			len("The name field must be between 2 to 20 characters."),
		),
		validation.Field(&u.Email, vre("The Email  is required")),
		validation.Field(&u.Password, vre("The Password  is required")),
	)
}

// type UserList struct {
// 	Users []User `db:"users"`
// }

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
	pareseRegTemplate(w, User{
		CSRFToken: nosurf.Token(r),
	})
}

func (c connection) StoreUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	users := User{}

	if err := c.formDecoder.Decode(&users, r.PostForm); err != nil {
		log.Fatal(err)
	}

	if err := users.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			fmt.Println(vErr)
			users.FormError = vErr
		}
		pareseRegTemplate(w, users)
		return
	}

	log.Println(r.PostForm, users)

	createUserQuery := `
	INSERT INTO users(
		name,
		email,
		password
		)  VALUES(
			:name,
		:email,
		:password
		)
		returning *`

	stmt, _ := c.db.PrepareNamed(createUserQuery)

	stmt.Get(&users, users)

	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
