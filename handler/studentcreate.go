package handler

import (
	"html/template"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/justinas/nosurf"
)

type Student struct {
	ID        int       `db:"id" form:"-"`
	FirstName string    `db:"first_name" form:"first_name"`
	LastName  string    `db:"last_name" form:"last_name"`
	Class     string    `db:"class" form:"class"`
	Roll      int       `db:"roll" form:"roll"`
	Email     string    `db:"email" form:"email"`
	Password  string    `db:"password" form:"password"`
	CreatedAt time.Time `db:"created_at" form:"created_at"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at"`
	CSRFToken string    `db:"-" form:"csrf_token"`
	FormError map[string]error
}

func (s Student) ValidateStudent() error {
	vre := validation.Required.Error
	len := validation.Length(2, 20).Error
	return validation.ValidateStruct(&s,
		validation.Field(&s.FirstName,
			vre("The FirstName  is required"),
			len("The FirstName field must be between 2 to 20 characters."),
		),
		validation.Field(&s.LastName,
			vre("The LastName  is required"),
			len("The LastName field must be between 2 to 20 characters."),
		),
		validation.Field(&s.Class, vre("The Class  is required")),
		validation.Field(&s.Roll, vre("The Roll  is required")),
		validation.Field(&s.Email, vre("The Email  is required")),
		validation.Field(&s.Password, vre("The Password  is required")),
	)
}

func pareseStudentTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/createstudent.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err := t.ExecuteTemplate(w, "createstudent.html", data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) CreateStudent(w http.ResponseWriter, r *http.Request) {
	pareseStudentTemplate(w, Student{
		CSRFToken: nosurf.Token(r),
	})
}

func (c *connection) StoreStudent(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	students := Student{}

	if err := c.formDecoder.Decode(&students, r.PostForm); err != nil {
		log.Fatal(err)
	}

	if err := students.ValidateStudent(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			students.FormError = vErr
		}
		pareseStudentTemplate(w, students)
		return
	}

	log.Println(r.PostForm, students)

	createStudentQuery := `
	INSERT INTO students (
	 first_name, 
	 last_name, 
	 class,
	 roll, 
	 email, 
	 password
	 ) VALUES (
		:first_name, 
		:last_name, 
		:class, 
		:roll, 
		:email, 
		:password
		)
		returning *`

	stmt, _ := c.db.PrepareNamed(createStudentQuery)

	stmt.Get(&students, students)

	http.Redirect(w, r, "/list/student", http.StatusSeeOther)
}
