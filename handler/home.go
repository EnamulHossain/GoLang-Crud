package handler

import (
	
	"net/http"
)


func (h connection) Home(w http.ResponseWriter, r *http.Request) {
	h.pareseHomeTemplate(w, r)
}




func (h connection) LHome(w http.ResponseWriter, r *http.Request) {
	h.pareseLHomeTemplate(w, r)
}

