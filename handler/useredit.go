package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
	// "github.com/go-chi/chi/v5"
)

func pareseEditUserTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/edituser.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.ExecuteTemplate(w, "edituser.html", data)
}

func (c connection) EditUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var UserEdit User

	const editUserQuery = `Select * FROM users WHERE id = $1`
	if err := c.db.Get(&UserEdit, editUserQuery, id); err != nil {
		log.Fatalln(err)
	}

	UserEdit.CSRFToken = nosurf.Token(r)
	pareseEditUserTemplate(w, UserEdit)

}

func (c connection) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatalln("%#V", err)
	}

	user := User{ID: uID}
	if err := c.formDecoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}

	const UpdateQQ = `
	UPDATE Users SET
	    name =:name,
		email = :email,
		password = :password
		WHERE id= :id;
	`

	stmt, err := c.db.PrepareNamed(UpdateQQ)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := stmt.Exec(user)
	if err != nil {
		log.Fatalln(err)
	}
	Rcount, err := res.RowsAffected()
	if Rcount < 1 || err != nil {
		log.Fatalln(err)
	}
	http.Redirect(w, r, "/user/list", http.StatusSeeOther)

}
