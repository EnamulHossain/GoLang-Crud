package handler

import (
	"net/http"
	"strings"

)

func (c *connection) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	
	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/student/delete/", "")

	
	if  err := c.storage.DeleteStudentByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list/student", http.StatusPermanentRedirect)
}
