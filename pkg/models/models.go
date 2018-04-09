package models

import (
	"time"
)

type Project struct {
	ID int
	Title string
	Data string
	Created time.Time
	Tags int
	Views int
	Tagline string
	Authors int
}

type Author struct {
	ID int
	FirstName string
	LastName string
	Bio string
	Projects int
}

type Projects []*Project
