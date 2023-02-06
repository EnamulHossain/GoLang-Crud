package main

import (
	"StudentManagement/handler"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var sessionManager *scs.SessionManager

func main() {

	// Start Enviroment
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
	// End Enviroment

	// Start  Sesson
	lt := config.GetDuration("session.lifetime")
	it := config.GetDuration("session.idletime")
	sessionManager = scs.New()
	sessionManager.Lifetime = lt * time.Hour
	sessionManager.IdleTimeout = it * time.Minute
	sessionManager.Cookie.Name = "web-session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true
	// End Sesson

	// Start Database Connection
	db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=studentmanagement sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	// End Database Connection
	userMigration := `
    CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`

	studentMigration := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		class VARCHAR(255) NOT NULL,
		roll INT NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	  );
	  `

	markMigration := `
	CREATE TABLE IF NOT EXISTS marks (
		id SERIAL PRIMARY KEY,
		student_id INT REFERENCES students(id) ON DELETE CASCADE,
		datastructures INT NOT NULL,
		algorithms INT NOT NULL,
		computernetworks INT NOT NULL,
		artificialintelligence INT NOT NULL,
		operatingsystems INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	  );`

	db.MustExec(userMigration)
	db.MustExec(studentMigration)
	db.MustExec(markMigration)

	p := config.GetInt("server.port")

	decoder := form.NewDecoder()

	_, chi := handler.New(db, sessionManager, decoder)

	newChi := nosurf.New(chi)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), sessionManager.LoadAndSave(newChi)); err != nil {
		log.Fatalf("%#v", err)
	}
}
