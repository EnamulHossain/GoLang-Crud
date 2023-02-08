package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type PostgresStorage struct {
	DB *sqlx.DB
}

func NewPostgresStorage (config *viper.Viper) (*PostgresStorage,error) {
	db,err := ConnectDatabase(config);
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		DB: db,
	},nil
}


func ConnectDatabase(config *viper.Viper) (*sqlx.DB,error) {

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
		return nil, err
	}

	return db,nil
}