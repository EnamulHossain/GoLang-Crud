package main

import (
	"StudentManagement/handler"
	"StudentManagement/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

var sessionManager *scs.SessionManager

// var migration = `
//     CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		name VARCHAR(255) NOT NULL,
// 		email VARCHAR(255) UNIQUE NOT NULL,
// 		password VARCHAR(255) NOT NULL,
// 		status BOOLEAN DEFAULT TRUE,
// 		created_at TIMESTAMP DEFAULT NOW(),
// 		updated_at TIMESTAMP DEFAULT NOW(),
// 		deleted_at TIMESTAMP DEFAULT NULL
// 	);


// 	CREATE TABLE IF NOT EXISTS students (
// 		id SERIAL PRIMARY KEY,
// 		first_name VARCHAR(255) NOT NULL,
// 		last_name VARCHAR(255) NOT NULL,
// 		class VARCHAR(255) NOT NULL,
// 		roll INT NOT NULL,
// 		email VARCHAR(255) UNIQUE NOT NULL,
// 		password VARCHAR(255) NOT NULL,
// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		deleted_at TIMESTAMP DEFAULT NULL

// 	  );
	
// 	CREATE TABLE IF NOT EXISTS marks (
// 		id SERIAL PRIMARY KEY,
// 		student_id INT REFERENCES students(id) ON DELETE CASCADE,
// 		datastructures INT NOT NULL,
// 		algorithms INT NOT NULL,
// 		computernetworks INT NOT NULL,
// 		artificialintelligence INT NOT NULL,
// 		operatingsystems INT NOT NULL,
// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		deleted_at TIMESTAMP DEFAULT NULL
// 	  );
	  
// 	CREATE TABLE IF NOT EXISTS sessions (
// 		token TEXT PRIMARY KEY,
// 		data BYTEA NOT NULL,
// 		expiry TIMESTAMPTZ NOT NULL
// 		);	

// 	CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);
// `

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

	decoder := form.NewDecoder()

	postgresStorage, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}


	if err := goose.SetDialect("postgres"); err != nil {
        panic(err)
    }

    if err := goose.Up(postgresStorage.DB.DB, "migrations"); err != nil {
        panic(err)
    }

	// End Database Connection
	// Start  Sesson
	lt := config.GetDuration("session.lifetime")
	it := config.GetDuration("session.idletime")
	sessionManager = scs.New()
	sessionManager.Lifetime = lt * time.Hour
	sessionManager.IdleTimeout = it * time.Minute
	sessionManager.Cookie.Name = "web-session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true
	sessionManager.Store = NewSQLXStore(postgresStorage.DB)

	// End Sesson

	p := config.GetInt("server.port")

	_, chi := handler.New(postgresStorage, sessionManager, decoder)

	newChi := nosurf.New(chi)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), sessionManager.LoadAndSave(newChi)); err != nil {
		log.Fatalf("%#v", err)
	}
}
