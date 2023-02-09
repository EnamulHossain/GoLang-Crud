package handler

import (
	"net/http"
	"strings"

)

func (c *connection) DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/user/delete/", "")


	if err := c.storage.DeleteUserByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
