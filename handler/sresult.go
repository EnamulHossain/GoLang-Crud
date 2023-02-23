package handler

import (
	// "StudentManagement/storage"
	"StudentManagement/storage"
	"net/http"

)

type StsForm struct{
	Selectedstudentsub []storage.Result
	SelectedSubMark storage.Result
}


func (c connection) Result(w http.ResponseWriter, r *http.Request) {
	

	res,err:=c.storage.Resul()

	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	// data := StsForm{
	// 	Selectedstudentsub: res,
	// 	SelectedSubMark:    res[0],
	// }
	c.pareseResultTemplate(w,res)
}