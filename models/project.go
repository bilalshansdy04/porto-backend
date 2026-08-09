package models

import "time"

type Project struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Status       string    `json:"status"` // e.g. "Live" or "Draft"
	ImageURL     string    `json:"image_url"`
	DateModified time.Time `json:"date_modified" gorm:"autoUpdateTime"`

	// Using GORM's serializer to store JSON arrays as text in SQLite
	TechStack   []string `json:"tech_stack" gorm:"serializer:json"`
	ProjectFlow []string `json:"project_flow" gorm:"serializer:json"`
	JobDesc     []string `json:"jobdesc" gorm:"serializer:json"`
}
