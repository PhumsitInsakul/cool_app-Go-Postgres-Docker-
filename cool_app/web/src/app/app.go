package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/pharick/cool_app/db"
)

type App struct {
	router    *http.ServeMux
	templates *Templates
	db        *sql.DB
}

func NewApp() (*App, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	dbConn, err := db.Connect(&db.Settings{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		return nil, err
	}

	app := &App{
		router:    http.NewServeMux(),
		templates: NewTemplates(),
		db:        dbConn,
	}
	return app, nil
}

func (a *App) RegisterHandler(url string, handler http.HandlerFunc) {
	a.router.HandleFunc(url, handler)
}

func (a *App) Serve(port int) {
	log.Printf("Server started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), a.router)
}
