package handler

import (
	"StudentManagement/storage"
	"net/http"
)

type Stlist struct {
	Students []storage.Student
	SearchTerm string
}


func (c connection) ListStudent(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm();

	stt := r.FormValue("SearchTerm")
	sf := storage.StudentFilter{
		SearchTerm: stt,
	}

	listStudent,err:=c.storage.ListStudent(sf)

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	data := Stlist{
		Students: listStudent,
		SearchTerm: stt,
	}
	c.pareseStudentListTemplate(w, data)
}
