package handler

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func (c connection) CreateMark(w http.ResponseWriter, r *http.Request) {
	classList, err := c.storage.ListClass()
	if err != nil {
		log.Fatalln(err)
	}
	studentList, err := c.storage.ListStudent()
	if err != nil {
		log.Fatalln(err)
	}
	c.pareseMarkTemplate(w, UserForm{
		ClassList: classList,
		StudentList: studentList,
		CSRFToken: nosurf.Token(r),
	})
}
