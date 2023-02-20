package handler

import (
	"StudentManagement/storage"
	"net/http"
)

type Ulist struct {
	Users []storage.User
	SearchTerm string
}

func (c connection) UserList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();

	st := r.FormValue("SearchTerm")

	uf := storage.UserFilter{
		SearchTerm: st,
	}
	user, err := c.storage.ListUser(uf)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := Ulist{
		Users: user,
		SearchTerm: st,
	}

	c.pareseUserTemplate(w, data)
}
