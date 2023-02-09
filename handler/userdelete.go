package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (c *connection) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("##############")
	id := chi.URLParam(r, "id")

	log.Println("id====",id)

	if err := c.storage.DeleteUserByID(id); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/list", http.StatusSeeOther)
}
