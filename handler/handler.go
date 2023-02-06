package handler

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
	"github.com/jmoiron/sqlx"
)

type connection struct {
	db             *sqlx.DB
	sessionManager *scs.SessionManager
	formDecoder    *form.Decoder
}

func New(db *sqlx.DB, sm *scs.SessionManager, formDecoder *form.Decoder) (connection, *chi.Mux) {
	c := connection{
		db:             db,
		sessionManager: sm,
		formDecoder:    formDecoder,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", c.Home)
	r.Get("/login", c.Login)
	r.Get("/reg", c.Reg)

	r.Route("/user", func(r chi.Router) {

		r.Post("/store", c.StoreUser)
		r.Get("/list", c.UserList)
		r.Get("/delete/{{.ID}}", c.DeleteUser)
		r.Get("/{id:[0-9]+}/edit", c.EditUser)
		r.Post("/{id:[0-9]+}/update", c.UpdateUser)

	})

	r.Get("/create/student", c.CreateStudent)
	r.Post("/store/student", c.StoreStudent)
	r.Get("/list/student", c.ListStudent)
	r.Get("/student/delete/{{.ID}}", c.DeleteStudent)
	r.Get("/student/{id:[0-9]+}/edit", c.StudentEdit)
	r.Post("/student/{id:[0-9]+}/update", c.StudentUpdate)

	r.Post("/user/login", c.LoginUser)

	return c, r
}
