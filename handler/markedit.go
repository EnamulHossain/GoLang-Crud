package handler

import (
	"StudentManagement/storage"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
)
type MarkEdit struct{
	MarkIn storage.MarkEdit
	FormError map[string]error
	CSRFToken string
}

func (h connection) EditMark(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mark, err := h.storage.MarkEdit(id)

	if err != nil {
		log.Println(mark)
	}
	fmt.Printf("%#v", mark)
	h.PareseMarkeditTemplate(w, MarkEdit{
		MarkIn:    *mark,
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	})

}

func (h connection) UpdateMark(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err)
	}
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	var form MarkEdit
	user := storage.MarkEdit{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	form.MarkIn = user
	fmt.Printf("%#v",user)
	err1 := h.storage.UpdateMarksbyID(user.Marks,id)
	if err1 != nil {
        log.Println(err)
	}
	http.Redirect(w, r, "/student/list", http.StatusSeeOther)
}
