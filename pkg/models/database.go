package models

import (
	"database/sql"
)

type Database struct{
	*sql.DB
}

func (db *Database) GetProject(id int) (*Project, error) {
	stmt := `SELECT json, created, tagline, tags, views FROM projects
	WHERE id = ?`

	row := db.QueryRow(stmt, id)

	s := &Project{}

	err := row.Scan(&s.Data, &s.Created, &s.Tagline, &s.Tags, &s.Views)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (db *Database) LatestProjects() (Projects, error) {
	stmt := `SELECT title, tagline, coverPhoto FROM projects
			 ORDER BY created DESC LIMIT 10`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := Projects{}

	for rows.Next() {
		s := &Project{}
		err := rows.Scan(&s.Title, &s.Tagline, &s.CoverPhoto)
		if err != nil {
			return nil, err
		}
		projects = append(projects, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, err
}

func (db *Database) InsertProject(title, data, created, authors, tagline string) (int, error) {
	stmt := `INSERT INTO projects (title, json, created, authors, tagline)
	VALUES(?, ?, HST_TIMESTAMP(), ?, ?)`

	result, err := db.Exec(stmt, title, data, authors, tagline)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}