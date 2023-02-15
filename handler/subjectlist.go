package handler

import (
	"net/http"
)



func (c connection) ListSubject(w http.ResponseWriter, r *http.Request) {
	
	Listsubject,err:= c.storage.ListSubject()

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	c.pareseSujectListTemplate(w, Listsubject)
}
