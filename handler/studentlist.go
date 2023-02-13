package handler

import (
	"net/http"
)



func (c connection) ListStudent(w http.ResponseWriter, r *http.Request) {
	
	listStudent,err:=c.storage.ListStudent()

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	c.pareseStudentListTemplate(w, listStudent)
}
