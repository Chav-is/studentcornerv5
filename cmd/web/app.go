package main

import (
	"studentcorner.com/pkg/models"
)

type App struct {
	Database *models.Database
	HTMLDir string
	StaticDir string
}