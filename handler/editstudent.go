package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/justinas/nosurf"
	// "github.com/go-chi/chi/v5"
)

func pareseEditStudentTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/editstudent.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "editstudent.html", data)
}

func (c connection) StudentEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var StudentEdit Student

	const editStudentQuery = `Select * FROM students WHERE id = $1 AND deleted_at IS NULL`
	if err := c.db.Get(&StudentEdit, editStudentQuery, id); err != nil {
		log.Fatalln(err)
	}

	StudentEdit.CSRFToken = nosurf.Token(r)
	pareseEditStudentTemplate(w, StudentEdit)

}

func (h connection) StudentUpdate(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatalln("%#V", err)
	}
	// //............................................................
	student := Student{ID: uID}
	if err := h.formDecoder.Decode(&student, r.PostForm); err != nil {
		log.Fatal(err)
	}

	if err := student.ValidateStudent(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			student.FormError = vErr
		}
		pareseStudentTemplate(w, student)
		return
	}

	const UpdateQQ = `
	UPDATE Students SET
	    first_name =:first_name,
		last_name =:last_name,
		class = :class,
		roll = :roll,
		email = :email,
		password = :password
		WHERE id= :id;
	`
	//.....................................................................
	stmt, err := h.db.PrepareNamed(UpdateQQ)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := stmt.Exec(student)
	if err != nil {
		log.Fatalln(err)
	}
	Rcount, err := res.RowsAffected()
	if Rcount < 1 || err != nil {
		log.Fatalln(err)
	}

	http.Redirect(w, r, "/list/student", http.StatusSeeOther)

}
