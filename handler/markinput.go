package handler

import (
	"fmt"
	"net/http"
)





func (c connection) MarkInput(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		
	}

	var form MarkForm

	if err := c.decoder.Decode(&form, r.PostForm); err != nil {
		
	}
	fmt.Println("#######################################")
	fmt.Printf("%+v", form)
	fmt.Println("#######################################")
	alldata,_:= c.storage.GetMarkInputOptionByID(form.Student)
	c.pareseMarkinputTemplate(w, alldata)

}