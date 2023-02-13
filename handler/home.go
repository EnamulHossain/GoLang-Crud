package handler

import (
	"net/http"
)

func (c connection) Home(w http.ResponseWriter, r *http.Request) {
	c.pareseHomeTemplate(w, r)
}
