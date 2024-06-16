package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // if we don't have _ before string (link) go won't allow you to use it when pressed command + s
)

type Settings struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Connect(s *Settings) (*sql.DB, error) { // return object sql.DB and func can return error if something go wrong
	// create connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.User, s.Password, s.DBName,
	)

	// open database connection
	return sql.Open("postgres", connStr)
}
