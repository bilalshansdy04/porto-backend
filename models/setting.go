package models

type Setting struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	ShowExperience bool   `json:"show_experience" gorm:"default:true"`
	ShowPhoto      bool   `json:"show_photo" gorm:"default:true"`
	Language       string `json:"language" gorm:"default:'id'"`
}
