package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (c *connection) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	
	id := chi.URLParam(r,"id")
	
	if  err := c.storage.DeleteStudentByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list/student", http.StatusPermanentRedirect)
}
