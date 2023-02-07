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

var migration = `
    CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		status BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		deleted_at TIMESTAMP DEFAULT NULL
	);


	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		class VARCHAR(255) NOT NULL,
		roll INT NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL

	  );
	
	CREATE TABLE IF NOT EXISTS marks (
		id SERIAL PRIMARY KEY,
		student_id INT REFERENCES students(id) ON DELETE CASCADE,
		datastructures INT NOT NULL,
		algorithms INT NOT NULL,
		computernetworks INT NOT NULL,
		artificialintelligence INT NOT NULL,
		operatingsystems INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL
	  );
	  
	CREATE TABLE IF NOT EXISTS sessions (
		token TEXT PRIMARY KEY,
		data BYTEA NOT NULL,
		expiry TIMESTAMPTZ NOT NULL
		);	

	CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);
`

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

	decoder := form.NewDecoder()

	// Start Database Connection
	// db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=studentmanagement sslmode=disable")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.dbname"),
		config.GetString("database.sslmode"),
	))
	if err != nil {
		log.Fatalln(err)
	}
	// sessionManager.Store = NewSQLXStore(db)

	res := db.MustExec(migration)
	row, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	if row < 0 {
		log.Fatalln("failed to run schema")
	}
	// End Database Connection

	// db.MustExec(migration)

	p := config.GetInt("server.port")

	_, chi := handler.New(db, sessionManager, decoder)

	newChi := nosurf.New(chi)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), sessionManager.LoadAndSave(newChi)); err != nil {
		log.Fatalf("%#v", err)
	}
}
