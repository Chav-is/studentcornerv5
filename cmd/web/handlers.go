package main

import (
	"fmt"
	"net/http"
	"strconv"
	"studentcorner.com/pkg/forms"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	projects, err := app.Database.LatestProjects()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r,"home.page.html", &HTMLData{
		Projects: projects,
	})

}

func (app *App) ShowProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	project, err := app.Database.GetProject(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if project == nil {
		app.NotFound(w)
		return
	}

	app.RenderHTML(w, r,"show.page.html", &HTMLData{
		Project: project,
	})
}

func (app *App) NewProject(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "new.page.html", &HTMLData{
		Form: &forms.NewProject{},
	})
}

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.NewProject{
		Title: r.PostForm.Get("title"),
		Authors: r.PostForm.Get("authors"),
		Data: r.PostForm.Get("data"),
		Created: r.PostForm.Get("created"),
		Tagline: r.PostForm.Get("tagline"),

	}
	if !form.Valid() {
		app.RenderHTML(w, r, "new.page.html", &HTMLData{Form: form})
		return
	}
	id, err := app.Database.InsertProject(form.Title, form.Authors, form.Data, form.Created, form.Tagline)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/project/%d", id), http.StatusSeeOther)
}