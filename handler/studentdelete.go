package handler

import (
	"net/http"
	"strings"
)

func (c *connection) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/student/delete/", "")

	deleteStudentQuery := `
	DELETE FROM students where id = $1`

	res := c.db.MustExec(deleteStudentQuery, id)

	if ok, err := res.RowsAffected(); err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list/student", http.StatusPermanentRedirect)
}
