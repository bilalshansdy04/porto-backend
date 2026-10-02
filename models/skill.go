package models

import "gorm.io/gorm"

type Skill struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (p *Skill) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == 0 {
		var maxId int64
		tx.Model(&Skill{}).Select("IFNULL(MAX(id), 0)").Scan(&maxId)
		p.ID = uint(maxId + 1)
	}
	return
}
