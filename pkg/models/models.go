package models

import (
	"time"
)

type Project struct {
	ID int `json:"pid"`
	Title string `json:"title"`
	Data string `json:"Data"`
	Created time.Time `json:"Created"`
	Tags int `json:"tags"`
	Views int `json:"Views"`
	Tagline string `json:"Tagline"`
	Authors int `json:"authors"`
	CoverPhoto string `json:"CoverPhoto"`
}

type Author struct {
	ID int `json:"uid"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Bio string `json:"bio"`
	Projects int `json:"projects"`
}

type Projects []*Project
