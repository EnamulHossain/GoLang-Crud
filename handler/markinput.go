package handler

import (
	"StudentManagement/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)





func (c connection) MarkInput(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	var form MarkForm

	if err := c.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln(err)
		
	}
	fmt.Println("#######################################")
	fmt.Printf("%+v", form)
	fmt.Println("#######################################")
	alldata,_:= c.storage.GetMarkInputOptionByID(form.Student)
	// c.pareseMarkinputTemplate(w, alldata)

	c.pareseMarkinputTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
		MarkInput: alldata,
	})

}



func (c connection) StoreMark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	
	marks:=storage.StudentSubject{}
	
	if err := c.decoder.Decode(&marks, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println("show mark map",marks.Mark)	

        
		for id, mark := range marks.Mark {
		m := storage.StudentSubject{
			ID: id,
			Marks:     mark,
		}
		_, err := c.storage.Markcreate(m)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, "/mark/create", http.StatusSeeOther)
	

}

