package handler

import (
	"net/http"
)

// func pareseUserTemplate(w http.ResponseWriter, data any) {
// 	t, err := template.ParseFiles("./template/header.html", "./template/footer.html", "./template/userlist.html")
// 	if err != nil {
// 		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
// 	}
// 	t.ExecuteTemplate(w, "userlist.html", data)
// }

func (c connection) UserList(w http.ResponseWriter, r *http.Request) {
	user, err := c.storage.ListUser()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	c.pareseUserTemplate(w, user)
}
