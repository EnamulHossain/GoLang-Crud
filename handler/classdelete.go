package handler

import (
	"net/http"
	"strings"
)

func (c *connection) DeleteClass(w http.ResponseWriter, r *http.Request) {
	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/class/delete/", "")

	if err := c.storage.DeleteClassByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/class/list", http.StatusSeeOther)
}
