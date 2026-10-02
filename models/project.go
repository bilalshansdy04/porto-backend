package models

import "time"

type ProjectItem struct {
	Text      string `json:"text"`
	IsVisible bool   `json:"is_visible"`
}

type ProjectScreenshot struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type Project struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Status       string    `json:"status"` // e.g. "Live" or "Draft"
	ImageURL     string    `json:"image_url"`
	DateModified time.Time `json:"date_modified" gorm:"autoUpdateTime"`
	IsVisible    bool      `json:"is_visible"`

	// Using GORM's serializer to store JSON arrays as text
	TechStack      []string            `json:"tech_stack" gorm:"serializer:json"`
	ProjectFlow    []ProjectItem       `json:"project_flow" gorm:"serializer:json"`
	JobDesc        []ProjectItem       `json:"jobdesc" gorm:"serializer:json"`
	Link           string              `json:"link"`
	CarouselImages []string            `json:"carousel_images" gorm:"serializer:json"`
	Screenshots    []ProjectScreenshot `json:"screenshots" gorm:"serializer:json"`
}
