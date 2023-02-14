package handler

import (
	"log"
	"net/http"
)

func (c connection) pareseRegTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("reg.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}




func (c connection) pareseLoginTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("login.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	t.Execute(w,  data)
}








func(c connection) pareseEditUserTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("edituser.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
	
}



func (c connection) pareseUserTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("userlist.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Println(err)
	}
}




//Student




func(c connection) pareseStudentTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("createstudent.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}



func(c connection) pareseEditStudentTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("editstudent.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	t.Execute(w, data)
}


func (c connection) pareseStudentListTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("studentlist.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	t.Execute(w, data)
}




func(c connection) pareseHomeTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("home.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}