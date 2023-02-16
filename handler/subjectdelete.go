package handler

import (
	"net/http"
	"strings"

)

func (c *connection) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	
	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/subject/delete/", "")


	if err := c.storage.DeleteSubjectByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/subject/list", http.StatusSeeOther)
}
