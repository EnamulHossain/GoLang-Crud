package handler

import (
	"StudentManagement/storage"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
)

type connection struct {
	storage        dbStorage
	sessionManager *scs.SessionManager
	formDecoder    *form.Decoder
}
type dbStorage interface {
	ListUser() ([]storage.User, error)
	CreateUser(u storage.User) (*storage.User, error)
	UpdateUser(u storage.User) (*storage.User, error)
	GetUserByID(id string) (*storage.User, error)
	GetUserByUsername(username string) (*storage.User, error)
	DeleteUserByID(id string) error


	CreateStudent(u storage.Student) (*storage.Student, error)
	ListStudent() ([]storage.Student, error)
	UpdateStudent(u storage.Student) (*storage.Student, error)
	GetStudentByID(id string) (*storage.Student, error)
	GetStudentByUsername(username string) (*storage.Student, error)
	DeleteStudentByID(id string) error
}

func New(storage dbStorage, sm *scs.SessionManager, formDecoder *form.Decoder) (connection, *chi.Mux) {
	c := connection{
		sessionManager: sm,
		formDecoder:    formDecoder,
		storage:        storage ,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", c.Home)
	r.Get("/login", c.Login)
	r.Get("/reg", c.Reg)

	r.Get("/create/student", c.CreateStudent)
	r.Post("/student/store", c.StoreStudent)
	r.Get("/list/student", c.ListStudent)
	r.Get("/student/delete/{{.ID}}", c.DeleteStudent)
	r.Get("/student/{id:[0-9]+}/edit", c.StudentEdit)
	r.Post("/student/{id:[0-9]+}/update", c.StudentUpdate)

	r.Route("/user", func(r chi.Router) {

		r.Post("/store", c.StoreUser)
		r.Get("/list", c.UserList)
		r.Get("/delete/{{.ID}}", c.DeleteUser)
		r.Get("/{id:[0-9]+}/edit", c.EditUser)
		r.Post("/{id:[0-9]+}/update", c.UpdateUser)

	})

	r.Post("/user/login", c.LoginUser)

	return c, r
}
