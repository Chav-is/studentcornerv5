package forms

import(
	"strings"
	"unicode/utf8"
)

type NewProject struct {
	Title string
	Data string
	Created string
	Authors string
	Tagline string
	Failures map[string]string
}

func (f *NewProject) Valid() bool {
	f.Failures = make(map[string]string)

	if strings.TrimSpace(f.Title) == "" {
		f.Failures["Title"] = "Title is required"
	} else if utf8.RuneCountInString(f.Title) > 100 {
		f.Failures["Title"] = "Title cannot be longer than 100 characters"
	}

	if strings.TrimSpace(f.Data) == "" {
		f.Failures["Content"] = "Content is required"
	}

	if strings.TrimSpace(f.Authors) == "" {
		f.Failures["Authors"] = "Someone had to have made this"
	}

	if strings.TrimSpace(f.Tagline) == "" {
		f.Failures["Tagline"] = "You need a tag-line"
	}

	return len(f.Failures) == 0
}