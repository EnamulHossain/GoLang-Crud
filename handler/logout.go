package handler

import (
	"log"
	"net/http"
)

func (c connection) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := c.sessionManager.Destroy(r.Context()); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
