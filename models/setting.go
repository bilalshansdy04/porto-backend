package models

import "gorm.io/gorm"

type Setting struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	ShowExperience bool   `json:"show_experience" gorm:"default:true"`
	ShowPhoto      bool   `json:"show_photo" gorm:"default:true"`
	Language       string `json:"language" gorm:"default:'id'"`
}

func (p *Setting) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == 0 {
		var maxId int64
		tx.Model(&Setting{}).Select("IFNULL(MAX(id), 0)").Scan(&maxId)
		p.ID = uint(maxId + 1)
	}
	return
}
