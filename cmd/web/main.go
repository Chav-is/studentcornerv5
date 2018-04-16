package main

import (
	"log"
	"net/http"
	"flag"

	"studentcorner.com/pkg/models"

	_"github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP netowrk address")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	dsn := flag.String("dsn", "sc:@/studentcorner?parseTime=true", "MySQL DSN")

	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	app := &App{
		Database:  &models.Database{db},
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

 	http.Handle("/", http.FileServer(http.Dir(*htmlDir)))
	log.Printf("Serving %s on HTTP port: %s\n", *htmlDir, *addr)
	err := http.ListenAndServe(":4000", app.Routes())
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