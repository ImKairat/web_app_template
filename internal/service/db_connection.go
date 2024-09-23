package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"ses_back/internal/config"
)

var DB *sql.DB

func Connect() error {
	config.ReadConfig()
	psqlCfg := config.Config.Psql
	var err error

	DB, err = sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		psqlCfg.User, psqlCfg.Password, psqlCfg.Host, psqlCfg.Port, psqlCfg.DbName,
	))

	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Disconnect() error {
	err := DB.Close()
	if err != nil {

		return err
	}

	return nil
}
