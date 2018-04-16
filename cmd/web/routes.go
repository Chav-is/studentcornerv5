package main

import (
	"net/http"
	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	// mux.Post("/modifyproject/:id", http.HandlerFunc(app.EditProject))
	mux.Get("/snippet/new", http.HandlerFunc(app.NewProject))
	mux.Get("/project/:id", http.HandlerFunc(app.ShowProject))

	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

