package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func (c connection) CreateMark(w http.ResponseWriter, r *http.Request) {

	// Url := r.URL.Path
	// id := strings.ReplaceAll(Url, "/subject/delete/", "")
	
	classList, err := c.storage.ListClass()
	if err != nil {
		log.Fatalln(err)
	}
	studentList, err := c.storage.ListStudent(storage.StudentFilter{
		SearchTerm: "",
	})
	if err != nil {
		log.Fatalln(err)
	}
	c.pareseMarkTemplate(w, UserForm{
		ClassList: classList,
		StudentList: studentList,
		CSRFToken: nosurf.Token(r),
	})
}
