package job

import "time"

type Job struct {
	ID           string
	Title        string
	CompanyID    string
	City         string
	Category     string
	Experience   string
	Education    string
	Salary       string
	Status       string
	PublishedAt  time.Time
	UpdatedAt    time.Time
	SourceName   string
	SourceURL    string
	Description  string
	Requirements []string
	Skills       []string
}
