package models

import (
	"time"
	"gorm.io/gorm"
)

type ProjectItem struct {
	Text      string `json:"text"`
	IsVisible bool   `json:"is_visible"`
}

type ImageScreenshot struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProjectID uint   `json:"project_id"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
}

func (p *ImageScreenshot) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == 0 {
		var maxId int64
		tx.Model(&ImageScreenshot{}).Select("IFNULL(MAX(id), 0)").Scan(&maxId)
		p.ID = uint(maxId + 1)
	}
	return
}

type Project struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Status       string    `json:"status"` // e.g. "Live" or "Draft"
	ImageURL     string    `json:"image_url"`
	DateModified time.Time `json:"date_modified" gorm:"autoUpdateTime"`
	IsVisible    bool      `json:"is_visible"`

	// Using GORM's serializer to store JSON arrays as text
	TechStack      []string          `json:"tech_stack" gorm:"serializer:json"`
	ProjectFlow    []ProjectItem     `json:"project_flow" gorm:"serializer:json"`
	JobDesc        []ProjectItem     `json:"jobdesc" gorm:"serializer:json"`
	Link           string            `json:"link"`
	CarouselImages []string          `json:"carousel_images" gorm:"serializer:json"`
	Screenshots    []ImageScreenshot `json:"screenshots" gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == 0 {
		var maxId int64
		tx.Model(&Project{}).Select("IFNULL(MAX(id), 0)").Scan(&maxId)
		p.ID = uint(maxId + 1)
	}
	return
}
