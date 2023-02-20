package handler

import (
	"StudentManagement/storage"
	"net/http"
)

type Sulist struct {
	Subjects []storage.Subject
	SearchTerm string
}


func (c connection) ListSubject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	sut := r.FormValue("SearchTerm")
	
	suf := storage.SubjectFilter{
		SearchTerm: sut,
	}
	Listsubject,err:= c.storage.ListSubject(suf)

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	data := Sulist{
		Subjects: Listsubject,
		SearchTerm: sut,
	}
	c.pareseSujectListTemplate(w, data)
}
