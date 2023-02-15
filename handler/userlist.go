package handler

import (
	"net/http"
)



func (c connection) UserList(w http.ResponseWriter, r *http.Request) {
	user, err := c.storage.ListUser()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	c.pareseUserTemplate(w, user)
}
