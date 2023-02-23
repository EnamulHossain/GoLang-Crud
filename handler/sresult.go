package handler

import (
	// "StudentManagement/storage"
	"StudentManagement/storage"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type StsForm struct{
	Selectedstudentsub []storage.Result
	SelectedSubMark storage.Result
}


func (c connection) Result(w http.ResponseWriter, r *http.Request) {
	id:=chi.URLParam(r,"id")
	uid,_:=strconv.Atoi(id)

	res,err:=c.storage.Resul(uid)

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	data := StsForm{
		Selectedstudentsub: res,
		SelectedSubMark:    res[0],
	}
	c.pareseResultTemplate(w,data)
}




func (c connection) AllResult(w http.ResponseWriter, r *http.Request) {
	
	res,err:=c.storage.AllResult()

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	// data := StsForm{
		
	// 	Selectedstudentsub: res,
	// 	SelectedSubMark:    res[5],
	// }
	c.pareseAllResultTemplate(w,res)
}