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
	t.Execute(w, data)
}

func (c connection) pareseEditUserTemplate(w http.ResponseWriter, data any) {
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

//Student///////////////////////////////////////////////////////////////////////////////

func (c connection) pareseStudentTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("createstudent.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseEditStudentTemplate(w http.ResponseWriter, data any) {
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

func (c connection) pareseHomeTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("home.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

// class /////////////////////////////////////////////////////////////////

func (c connection) pareseClassTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("createclass.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseClassListTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("classlist.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseClassEditTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("editclass.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

//   Subject

func (c connection) pareseSubjectTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("subjectcreate.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseSujectListTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("subjectlist.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseEditSubjectTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("subjectedit.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	t.Execute(w, data)
}

// Mark

func (c connection) pareseMarkTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("addmark.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseLHomeTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("lhome.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (c connection) pareseMarkinputTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("input.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
