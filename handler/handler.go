package handler

import (
	"StudentManagement/storage"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
)

type connection struct {
	storage        dbStorage
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	Templates      *template.Template
}
type dbStorage interface {
	ListUser() ([]storage.User, error)
	CreateUser(u storage.User) (*storage.User, error)
	UpdateUser(u storage.User) (*storage.User, error)
	GetUserByID(id string) (*storage.User, error)
	GetUserByUsername(name string) (*storage.User, error)
	DeleteUserByID(id string) error

	CreateStudent(u storage.Student) (*storage.Student, error)
	ListStudent() ([]storage.Student, error)
	UpdateStudent(u storage.Student) (*storage.Student, error)
	GetStudentByID(id string) (*storage.Student, error)
	GetStudentByUsername(username string) (*storage.Student, error)
	DeleteStudentByID(id string) error

	ListClass() ([]storage.Class, error)
	CreateClass(c storage.Class) (*storage.Class, error)
	UpdateClass(c storage.Class) (*storage.Class, error)
	GetClassByID(id string) (*storage.Class, error)
	DeleteClassByID(id string) error

	ListSubject() ([]storage.Subject, error)
	CreateSubject(storage.Subject) (*storage.Subject, error)
	DeleteSubjectByID(id string) error
	GetSubjectByID(id string) (*storage.Subject, error)
	UpdateSubject(u storage.Subject) (*storage.Subject, error)



	GetSubjectByClassID(class int) ([]storage.Subject, error)
	InsertMark(s storage.StudentSubject) (*storage.StudentSubject, error)
}

func New(storage dbStorage, sm *scs.SessionManager, decoder *form.Decoder) (connection, *chi.Mux) {
	c := connection{
		sessionManager: sm,
		decoder:        decoder,
		storage:        storage,
	}

	c.ParseTemplates()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(Method)

	r.Group(func(r chi.Router) {
		// r.Use(sm.LoadAndSave)
		r.Get("/", c.Home)
		r.Get("/login", c.Login)
		r.Get("/reg", c.Reg)
		r.Post("/login", c.LoginPostHandler)
		r.Post("/user/store", c.StoreUser)

	})

	// workDir, _ := os.Getwd()
	// filesDir := http.Dir(filepath.Join(workDir, "assets/src"))
	// r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(c.Authentication)

		r.Route("/student", func(r chi.Router) {
			r.Get("/create", c.CreateStudent)
			r.Post("/store", c.StoreStudent)
			r.Get("/list", c.ListStudent)
			r.Get("/delete/{{.ID}}", c.DeleteStudent)
			r.Get("/{id:[0-9]+}/edit", c.StudentEdit)
			r.Post("/{id:[0-9]+}/update", c.StudentUpdate)
		})

		// SUBJECT
		r.Route("/subject", func(r chi.Router) {
			r.Get("/create", c.CreateSubject)
			r.Post("/store", c.StoreSubject)
			r.Get("/list", c.ListSubject)
			r.Get("/delete/{{.ID}}", c.DeleteSubject)
			r.Get("/{id:[0-9]+}/edit", c.EditSubject)
			r.Post("/{id:[0-9]+}/update", c.SubjectUpdate) //
		})

		// Class
		r.Route("/class", func(r chi.Router) {
			r.Get("/create", c.CreateClass)
			r.Post("/store", c.StoreClass)
			r.Get("/list", c.ListClass)
			r.Get("/delete/{{.ID}}", c.DeleteClass)
			r.Get("/{id:[0-9]+}/edit", c.ClassEdit)
			r.Post("/{id:[0-9]+}/update", c.ClassUpdate)
		})

		// Mark
		r.Route("/mark", func(r chi.Router) {
			r.Get("/create", c.CreateMark)
			// r.Post("/store", c.StoreClass)
			// r.Get("/list", c.ListClass)
			// r.Get("/delete/{{.ID}}", c.DeleteClass)
			// r.Get("/{id:[0-9]+}/edit", c.ClassEdit)
			// r.Post("/{id:[0-9]+}/update", c.ClassUpdate)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/list", c.UserList)
			r.Get("/delete/{{.ID}}", c.DeleteUser)
			r.Get("/{id:[0-9]+}/edit", c.EditUser)
			r.Post("/{id:[0-9]+}/update", c.UpdateUser)
		})
	})

	r.Get("/logout", c.LogoutHandler)

	return c, r
}

func Method(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch strings.ToLower(r.PostFormValue("_method")) {
			case "put":
				r.Method = http.MethodPut
			case "patch":
				r.Method = http.MethodPatch
			case "delete":
				r.Method = http.MethodDelete
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (h connection) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userName := h.sessionManager.GetString(r.Context(), "userName")
		userNamem := userName
		if userNamem == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *connection) ParseTemplates() error {
	templates := template.New("StudentManagement-template").Funcs(template.FuncMap{
		"globalfunc": func(n string) string {
			return ""
		},
	}).Funcs(sprig.FuncMap())
	newFS := os.DirFS("assets/template")
	tmpl := template.Must(templates.ParseFS(newFS, "*/*/*.html", "*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
