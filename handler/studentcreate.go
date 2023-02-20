package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)




func (c connection) CreateStudent(w http.ResponseWriter, r *http.Request) {
	classList,err := c.storage.ListClass()
	if err != nil {
		log.Fatalln(err)
	}

	c.pareseStudentTemplate(w, UserForm{
		ClassList: classList,
		CSRFToken: nosurf.Token(r),
	})
}

func (c *connection) StoreStudent(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	form :=UserForm{}
	students := storage.Student{}

	if err := c.decoder.Decode(&students, r.PostForm); err != nil {
		log.Fatal(err)
	}

	form.Student = students
	if err := students.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		c.pareseStudentTemplate(w, form)
		return
	}


	

	data,err:= c.storage.CreateStudent(students)
	if err != nil {
		log.Fatalln(err)
	}

	c.MarksHandler(w,r, students.Class, data.ID)

	http.Redirect(w, r, "/student/list", http.StatusSeeOther)
}
