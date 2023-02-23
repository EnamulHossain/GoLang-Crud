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
