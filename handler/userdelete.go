package handler

import (
	"net/http"
	"strings"
)

func (c *connection) DeleteUser(w http.ResponseWriter, r *http.Request) {
	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/user/delete/", "")

	deleteUserQuery := `
	DELETE FROM users where id = $1`

	res := c.db.MustExec(deleteUserQuery, id)

	if ok, err := res.RowsAffected(); err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/list", http.StatusPermanentRedirect)
}
