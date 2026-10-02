package models

import "gorm.io/gorm"

type Profile struct {
	ID                uint   `json:"id" gorm:"primary_key"`
	Summary           string `json:"summary" gorm:"type:varchar(200)"`
	YearsOfExperience int    `json:"years_of_experience"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	var maxId int64
	tx.Model(&Profile{}).Select("IFNULL(MAX(id), 0)").Scan(&maxId)
	p.ID = uint(maxId + 1)
	return
}
