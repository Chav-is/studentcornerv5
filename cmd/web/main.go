package main

import (
	"log"
	"net/http"
	"flag"

	"snippetbox.org/pkg/models"

	_"github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP netowrk address")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	dsn := flag.String("dsn", "sb:papaya77@/snippetbox?parseTime=true", "MySQL DSN")

	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	app := &App{
		Database:  &models.Database{db},
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}