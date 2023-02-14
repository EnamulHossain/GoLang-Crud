package handler

import (
	"net/http"
)



func (c connection) ListClass(w http.ResponseWriter, r *http.Request) {
	
	ListClass,err:=c.storage.ListClass()

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	c.pareseClassListTemplate(w, ListClass)
}
